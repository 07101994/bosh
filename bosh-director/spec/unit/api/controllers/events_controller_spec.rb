require 'spec_helper'
require 'rack/test'
require 'bosh/director/api/controllers/events_controller'

module Bosh::Director
  module Api
    describe Controllers::EventsController do
      include Rack::Test::Methods

      subject(:app) { described_class.new(config) }
      let(:config) { Config.load_hash(Psych.load(spec_asset('test-director-config.yml'))) }
      let(:temp_dir) { Dir.mktmpdir }
      let(:timestamp) { Time.now }

      before do
        App.new(config)
      end

      after { FileUtils.rm_rf(temp_dir) }

      context 'events' do
        before do
          Models::Event.make(
              "timestamp"    => timestamp,
              "user" => "test",
              "action"  => "create",
              "object_type"  => "deployment",
              "object_name" => "depl1",
              "task" => "1"
          )
          Models::Event.make(
              "parent_id"  => 1,
              "timestamp"    => timestamp,
              "user" => "test",
              "action"  => "create",
              "object_type"  => "deployment",
              "object_name" => "depl1",
              "task" => "2",
          )
        end

        it 'requires auth' do
          get '/'
          expect(last_response.status).to eq(401)
        end


        it 'returns a list of events' do
          basic_authorize 'admin', 'admin'
          get '/'

          expect(last_response.status).to eq(200)
          body = Yajl::Parser.parse(last_response.body)

          expect(body.size).to eq(2)

          expected = [
              {'id'          => '2',
               'parent_id'   => '1',
               'timestamp'   => timestamp.to_i,
               'user'        => 'test',
               'action'      => 'create',
               'object_type' => 'deployment',
               'object_name'   => 'depl1',
               'task'        => '2',
               'context'     => {}
              },
              {
                  'id'          => '1',
                  'timestamp'   => timestamp.to_i,
                  'user'        => 'test',
                  'action'      => 'create',
                  'object_type' => 'deployment',
                  'object_name'   => 'depl1',
                  'task'        => '1',
                  'context'     => {}
              }

          ]
          expect(Yajl::Parser.parse(last_response.body)).to eq(expected)
        end
      end
    end
  end
end
