## Director Console debugging

SSH to the Director VM

For director version <= 265.x
```
$ sudo su
# /var/vcap/jobs/director/bin/director_ctl console 
```

For director version >= 266.x
```
$ sudo su
# /var/vcap/jobs/director/bin/console
```

Potentially useful snippets:

```ruby
# list foreign keys for a table
Bosh::Director::Config.db.foreign_key_list(:instances)

# typical select
Bosh::Director::Config.db[:instances].where(variable_set_id: nil).first

# find first item in a table
Bosh::Director::Config.db[:schema_migrations].where(filename: "xxx").first

# consistent sort of applied migrations
Bosh::Director::Config.db[:schema_migrations].order(:filename).last

# manually record applied migration -- super dangerous.
Bosh::Director::Config.db[:schema_migrations] << {filename: "xxx"}

# update first item in a table
Bosh::Director::Models::Instance.first.update(boostrap: false)

# verify and delete disk reference
Bosh::Director::Models::PersistentDisk.where(disk_cid: "...").all
Bosh::Director::Models::PersistentDisk.where(disk_cid: "...").delete
```
