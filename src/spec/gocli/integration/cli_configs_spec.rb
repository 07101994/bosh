require_relative '../spec_helper'

describe 'cli configs', type: :integration do
  with_reset_sandbox_before_each

  let(:config) { yaml_file('config.yml', Bosh::Spec::NewDeployments.simple_cloud_config) }

  context 'can upload a config' do
    context 'when config uses placeholders' do
      let(:config) {yaml_file('config.yml', Bosh::Spec::Deployments.manifest_errand_with_placeholders)}

      it 'replaces placeholders' do
        expect(bosh_runner.run("update-config -v placeholder=my-data --type=my-type --name=default #{config.path}")).to include('Succeeded')
        expect(bosh_runner.run('config --type=my-type --name=default')).to include('my-data')
      end
    end

    it 'updates config' do
      expect(bosh_runner.run("update-config --type=my-type --name=default #{config.path}")).to include('Succeeded')
    end

    it 'updates named config' do
      expect(bosh_runner.run("update-config --type=my-type --name=my-name #{config.path}")).to include('Succeeded')
    end

    it 'updates config with default name' do
      bosh_runner.run("update-config --type=my-type --name=default #{config.path}")
      expect(bosh_runner.run('configs --type=my-type --json')).to include('"name": "default"')
    end

    it 'uploads an empty YAML hash' do
      Dir.mktmpdir do |tmpdir|
        empty_config_filename = File.join(tmpdir, 'empty_config.yml')
        File.write(empty_config_filename, '{}')
        expect(bosh_runner.run("update-config --type=my-type --name=default #{empty_config_filename}")).to include('Succeeded')
      end
    end

    it 'does not fail if the uploaded config is a large file' do
      config = Bosh::Common::DeepCopy.copy(Bosh::Spec::NewDeployments.simple_cloud_config)

      for i in 0..10001
        config["boshbosh#{i}"] = 'smurfsAreBlueGargamelIsBrownPinkpantherIsPinkAndPikachuIsYellow'
      end

      cloud_config_file = yaml_file('config.yml', config)

      output, exit_code = bosh_runner.run("update-config --type=large-config --name=default #{cloud_config_file.path}", return_exit_code: true)
      expect(output).to include('Succeeded')
      expect(exit_code).to eq(0)
    end
  end

  context 'can list configs' do
    let(:second_config) { yaml_file('second_config.yml', Bosh::Spec::Deployments.manifest_errand_with_placeholders) }

    it 'lists configs' do
      bosh_runner.run("update-config --type=my-type --name=default #{config.path}")
      bosh_runner.run("update-config --type=other-type --name=other-name #{config.path}")

      expect(bosh_runner.run('configs')).to include('default', 'other-name', 'my-type', 'other-type')
    end

    it 'can filter lists configs' do
      bosh_runner.run("update-config --type=my-type --name=my-name #{config.path}")
      bosh_runner.run("update-config --type=other-type --name=other-name #{config.path}")

      output = bosh_runner.run('configs --type=my-type --name=my-name')
      expect(output).to_not include('other-type', 'other-name')
    end

    it 'can include outdated configs' do
      bosh_runner.run("update-config --type=my-type --name=my-name #{config.path}")
      bosh_runner.run("update-config --type=my-type --name=my-name #{second_config.path}")

      output = bosh_runner.run('configs --recent=99')

      expect(output.scan('my-type').length).to be(2)
      expect(output.scan('my-name').length).to be(2)
    end
  end

  context 'when teams are used' do
    with_reset_sandbox_before_each(user_authentication: 'uaa')

    let(:production_env) {{'BOSH_CLIENT' => 'production_team', 'BOSH_CLIENT_SECRET' => 'secret'}}
    let(:admin_env) {{'BOSH_CLIENT' => 'test', 'BOSH_CLIENT_SECRET' => 'secret'}}
    let(:team_read_env){{'BOSH_CLIENT' => 'team-client-read-access', 'BOSH_CLIENT_SECRET' => 'team-secret'}}
    let(:team_admin_env){{'BOSH_CLIENT' => 'team-client', 'BOSH_CLIENT_SECRET' => 'team-secret'}}

    it 'shows configs of the same team only' do
      bosh_runner.run(
        "update-config --type=cloud --name=prod #{config.path}",
        client: production_env['BOSH_CLIENT'],
        client_secret: production_env['BOSH_CLIENT_SECRET'])

      bosh_runner.run(
        "update-config --type=cloud --name=team #{config.path}",
        client: team_admin_env['BOSH_CLIENT'],
        client_secret: team_admin_env['BOSH_CLIENT_SECRET']
      )

      production_configs = table(bosh_runner.run('configs', json: true, client: production_env['BOSH_CLIENT'],
        client_secret: production_env['BOSH_CLIENT_SECRET']))
      expect(production_configs.length).to eq(1)
      expect(production_configs.first).to include('name'=>'prod', 'team'=>'production_team', 'type'=>'cloud')

      team_configs = table(bosh_runner.run('configs', json: true, client: team_read_env['BOSH_CLIENT'],
        client_secret: team_read_env['BOSH_CLIENT_SECRET']))
      expect(team_configs.length).to eq(1)
      expect(team_configs.first).to include('name' => 'team', 'team' => 'ateam', 'type' => 'cloud')
    end

    it 'allows to create/delete team only for admin or team admin' do
      output = bosh_runner.run(
        "update-config --type=team-type --name=default #{config.path}",
        failure_expected: true,
        client: team_read_env['BOSH_CLIENT'],
        client_secret: team_read_env['BOSH_CLIENT_SECRET'],
      )
      expect(output).to include(`Director responded with non-successful status code '401' response '{"code":600000,"description":"Require one of the scopes: bosh.admin, bosh.deadbeef.admin"}'`)

      bosh_runner.run(
        "update-config --type=team-type --name=team-name1 #{config.path}",
        client: team_admin_env['BOSH_CLIENT'],
        client_secret: team_admin_env['BOSH_CLIENT_SECRET'],
      )
      bosh_runner.run(
        "update-config --type=team-type --name=team-name2 #{config.path}",
        client: team_admin_env['BOSH_CLIENT'],
        client_secret: team_admin_env['BOSH_CLIENT_SECRET'],
      )

      admin_configs = table(bosh_runner.run('configs', json: true, client: admin_env['BOSH_CLIENT'],
                                                 client_secret: admin_env['BOSH_CLIENT_SECRET']))
      expect(admin_configs.length).to eq(2)

      expect(admin_configs.first).to include('name' => 'team-name1', 'team' => 'ateam', 'type' => 'team-type')
      expect(admin_configs.last).to include('name' => 'team-name2', 'team' => 'ateam', 'type' => 'team-type')
    end

    it 'allows to create/delete team only for admin or team admin' do
      output = bosh_runner.run("update-config --type=team-type --name=default #{config.path}", failure_expected: true, client: team_read_env['BOSH_CLIENT'], client_secret: team_read_env['BOSH_CLIENT_SECRET'])
      expect(output).to include(`Director responded with non-successful status code '401' response '{"code":600000,"description":"Require one of the scopes: bosh.admin, bosh.deadbeef.admin"}'`)

      bosh_runner.run("update-config --type=team-type --name=team-name1 #{config.path}", client: team_admin_env['BOSH_CLIENT'], client_secret: team_admin_env['BOSH_CLIENT_SECRET'])
      bosh_runner.run("update-config --type=team-type --name=team-name2 #{config.path}", client: team_admin_env['BOSH_CLIENT'], client_secret: team_admin_env['BOSH_CLIENT_SECRET'])

      admin_configs = table(bosh_runner.run('configs', json: true, client: team_admin_env['BOSH_CLIENT'], client_secret: team_admin_env['BOSH_CLIENT_SECRET']))
      expect(admin_configs.length).to eq(2)

      expect(admin_configs.first).to include('name' => 'team-name1', 'team' => 'ateam', 'type' => 'team-type')
      expect(admin_configs.last).to include('name' => 'team-name2', 'team' => 'ateam', 'type' => 'team-type')

      output = bosh_runner.run(
        'delete-config --type=team-type --name=team-name1',
        failure_expected: true,
        client: team_read_env['BOSH_CLIENT'],
        client_secret: team_read_env['BOSH_CLIENT_SECRET'],
      )
      expect(output).to include('Require one of the scopes: bosh.admin, bosh.deadbeef.admin')

      expect(bosh_runner.run('delete-config --type=team-type --name=team-name1', client: admin_env['BOSH_CLIENT'], client_secret: admin_env['BOSH_CLIENT_SECRET'])).to include('Succeeded')
      expect(bosh_runner.run('delete-config --type=team-type --name=team-name2', client: team_admin_env['BOSH_CLIENT'], client_secret: team_admin_env['BOSH_CLIENT_SECRET'])).to include('Succeeded')

      admin_configs = table(bosh_runner.run('configs', json: true, client: admin_env['BOSH_CLIENT'], client_secret: admin_env['BOSH_CLIENT_SECRET']))
      expect(admin_configs.length).to eq(0)
    end
  end

  context 'can delete a config' do
    it 'delete a config' do
      bosh_runner.run("update-config --type=my-type --name=my-name #{config.path}")
      bosh_runner.run("update-config --type=other-type --name=other-name #{config.path}")

      expect(bosh_runner.run('delete-config --type=my-type --name=my-name')).to include('Succeeded')
      output = bosh_runner.run('configs')
      expect(output).to_not include('my-type','my-name')
      expect(output).to include('other-type', 'other-name')
    end

    it 'warns if there is nothing to delete' do
      output = bosh_runner.run('delete-config --type=my-type --name=potato')
      expect(output).to include('Succeeded')
      expect(output).to include('No configs to delete')
    end
  end

  it 'gives nice errors for common problems when uploading', no_reset: true do
    # not logged in
    expect(
      bosh_runner.run(
        "update-config --type=my-type --name=default #{config.path}",
        include_credentials: false,
        failure_expected: true,
      ),
    ).to include("Director responded with non-successful status code '401' response 'Not authorized: '")

    # no file
    expect(
      bosh_runner.run(
        'update-config --type=my-type --name=default /some/nonsense/file',
        failure_expected: true,
      ),
    ).to include('no such file or directory')

    # file not yaml
    Dir.mktmpdir do |tmpdir|
      config_filename = File.join(tmpdir, 'config.yml')
      File.write(config_filename, "---\n}}}invalid yaml!")
      expect(
        bosh_runner.run(
          "update-config --type=my-type --name=default #{config_filename}",
          failure_expected: true,
        ),
      ).to include('did not find expected node content')
    end
  end
end
