require 'json'
require 'pry'
require 'set'
require 'awesome_print'

raw = JSON.parse('{"test":{"stage_3":{"attempts":1,"passed":3,"first_try":"2015-04-14T17:25:02.535818822+08:00","first_pass":"2015-04-14T17:25:02.535818822+08:00","stage_1":{"attempts":1,"passed":3,"first_try":"2015-04-14T17:25:02.535818822+08:00","first_pass":"2015-04-14T17:25:02.535818822+08:00","latest_try":"2015-04-14T17:25:02.535818822+08:00"},"stage_2":{"attempts":1,"passed":1,"first_try":"2015-04-14T17:23:27.119841956+08:00","first_pass":"2015-04-14T17:23:27.119841956+08:00","latest_try":"2015-04-14T17:23:27.119841956+08:00"}}},"zoltan":{"stage_1":{"attempts":1,"passed":3,"first_try":"2015-04-14T17:25:02.535818822+08:00","first_pass":"2015-04-14T17:25:02.535818822+08:00","latest_try":"2015-04-14T17:25:02.535818822+08:00"},"stage_3":{"attempts":1,"passed":1,"first_try":"2015-04-14T17:23:27.119841956+08:00","first_pass":"2015-04-14T17:23:27.119841956+08:00","latest_try":"2015-04-14T17:23:27.119841956+08:00"}},"mark":{"stage_1":{"attempts":1,"passed":3,"first_try":"2015-04-14T17:25:02.535818822+08:00","first_pass":"2015-04-14T17:25:02.535818822+08:00","latest_try":"2015-04-14T17:25:02.535818822+08:00"},"stage_2":{"attempts":1,"passed":1,"first_try":"2015-04-14T17:23:27.119841956+08:00","first_pass":"2015-04-14T17:23:27.119841956+08:00","latest_try":"2015-04-14T17:23:27.119841956+08:00"}}}')


passed_stage3 = Set.new(raw.map { |k, v| v['stage_3'] != nil ? k : nil }.uniq - [nil])
passed_stage2 = Set.new(raw.map { |k, v| v['stage_2'] != nil ? k : nil }.uniq - [nil])
passed_stage1 = Set.new(raw.map { |k, v| v['stage_1'] != nil ? k : nil }.uniq - [nil])


puts 'stage 1 & 2 & 3'
ap passed_stage1 && passed_stage2 && passed_stage3
puts 'stage 1 & 2'
ap passed_stage1 && passed_stage2
puts 'stage 1'
ap passed_stage1
