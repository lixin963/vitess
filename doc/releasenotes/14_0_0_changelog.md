# Changelog of Vitess v14.0.0

### Announcement
#### CLI
* deprecation of CreateKeyspace flags #10488
#### General
* v14 release notes: Online DDL is production ready #10310
### Bug fixes
#### Backup and Restore
* Skip repl config in restore when active reparents disabled #9675
#### Build/CI
* Permission fixes for Docker builds #9619
* Fix output issue in the do_release script #9621
* VReplication V2 test: bypass newly added lag test which is causing flakiness in CI #9727
* Repair paths-filter workflows #10087
* Make go version check reliable by using double brackets #10125
#### Cluster management
* [grpcshim] Fix select race and move to internal package #9116
* Use AllPrivs user for table GC transitions #9690
* TabletManager: Check for context cancellation in loop within ChangeTabletType() #9842
* [topo/helpers] Make CopyShardReplications idempotent #9849
* Fix the race between PromoteReplica and replication manager tick #9859
* Propagate error from 'updateLocked' to client #10181
* Revert semi_sync change to reintroduce enable_semi_sync and remove durability_policy #10201
* The flag --online_ddl_check_interval belongs in vtctld #10320
* vt/wrangler: fix deleting primary tablet if `Shard` does not exist #10373
* topo: cache CellInfo along with conn #10408
#### Examples
* demo.go: fix call to ToBytes #9615
* Use latest vitess-operator example CRDs #9751
#### General
* Sanitize tx serializer log & error messages #9802
* Changes for make tools #10117
* Sanitize log messages #10367
* Rename error sanitizer file to follow naming convention #10406
#### Query Serving
* evalengine: numbers #9623
* Extract collation data to enable distinct aggregation #9639
* Add QueryRowsAffected and QueryRowsReturned to vttablet metrics and deprecate QueryRowCounts #9656
* Use the gen4 planner for queries with outer joins #9657
* Gen4: make sure to not merge unsharded tables from different keyspaces #9665
* Parser fix for CREATE VIEW statement #9693
* Add collation parsing to generated columns in Create Table #9694
* Fixed missing order/group by, limit, having in derived #9701
* Online DDL: resubmitting a migration with same UUID retries it in case it was cancelled or failed. #9704
* Query plan/rules: apply for multi-table statement #9747
* Add support for uint32 bind var + test #9757
* Stop rewriting `JoinCondition` with `USING` #9767
* Manage MySQL Replication Status States Properly #9853
* Fix `__sq_has_values1` error with `PulloutSubquery` #9855
* OnlineDDL executor: route all VReplicationExec through single function #9861
* fix: planner panic on derived tables sorting in query builder  #9869
* feat: add weightstring for distinct #9874
* ORDER BY scoping rules should not be used in group_concat #9939
* Fix: Sequence query to ignore reserved and transaction #9968
* fix: dual query with exists clause having system table query in it #9969
* fix: make concatenate and limit concurrent safe #9979
* ApplySchema: allow-zero-in-date embedded as query comment in call to ExecuteFetchAsDba #9998
* Fix: reserved connection retry logic when vttablet or mysql is down #10005
* gen4: Fix sub query planning when the outer query is a dual query #10007
* Fix parsing of bind variables in a few places #10015
* Route explain tab plan to the proper Keyspace #10027
* Re-add sanitize_log_messages flag and add tests for flags #10049
* OnlineDDL: double statement validation upon submission #10065
* Fix Gen4 only_full_group_by regression #10069
* Fix Gen4 group/order by with derived table #10074
* Only start SQL thread temporarily to WaitForPosition if needed #10104
* Vitess online ddl: modify a column from textual to non-textual, ignore charset #10116
* Fix StreamExecute in Gen4CompareV3 #10122
* Do not mutate replication state in WaitSourcePos and ignore tablets with SQL_Thread stopped in ERS #10148
* sqlparser: fix canonical formatting for enums #10149
* Fix for empty results when no shards can be found to route to #10152
* fix: handle reserved connection reset when tx killer has locked the connection #10153
* Do not send field query when using reserved connection #10163
* Backwards compatible replication status to state transition #10167
* sqlparser: Handle case sensitive AST option for table options #10191
* Emit the ENGINE field for table options as case sensitive #10197
* check for connectionID before adding to querylist #10212
* Fix handling of unsigned and zerofill in parser and normalization #10220
* Fix parsing of the foreign key constraint actions in different order #10224
* Fix handling of VISIBLE or INVISIBLE keyword on indexes #10243
* plancache: Lazy sysvar planner functions #10248
* evalengine: proper float to integer conversions #10252
* Fix formatting for function expressions and booleans #10255
* fix: concatenate engine primitive #10257
* fix: copy bindvariables in logstats #10260
* Move all schemadiff comparisons to canonical form #10261
* Do not cache plans that are invalid because of `--no_scatter` #10279
* collations: map utf8 to utf8mb3 #10287
* Fix parsing the special case convert charset logic #10288
* Deprecate flag --online_ddl_check_interval #10308
* Fix schema tracking issue when `PRIMARY` tablet changes #10335
* Remove normalization of integral length types #10336
* Fix failure when Unowned lookup IS NULL #10346
* Fix for empty results when no shards can be found to route to [v3] #10360
* Revert super_read_only config file changes #10366
* OnlineDDL/vitess fix: convert data type to JSON  #10390
* OnlineDDL/vrepl: more error codes, spatial types #10394
* Online DDL fix: instant is possible for VIRTUAL, not STORED #10411
* schemadiff: column name check in CHECK constraint is case insensitive #10413
* making log sanitization more precise #10417
* v3: Fix issue when no routes are found from a lookup #10422
* fix: use vcursor to execute the projection input #10435
* Tablet throttler: dynamic ThrottleMetricThreshold #10439
* Change use_super_read_only default back to false #10448
* fix: scalar aggregation engine primitive #10465
* fix: aggregation empty row on join with grouping and aggregations #10480
* Backport to release-14: Fix parsing of CAST() statements #10512 and #10514 #10517
* Fix casing of vitess migration syntax and comments printing #10535
* VReplication: more unrecoverable error codes #10559
* [14.0] fix: handle planner_version and planner-version correctly #10604
#### VReplication
* Update error description for deprecated command #9622
* Support VDiff across DB versions #9679
* VReplication Workflows: Use WithDDL while updating time_heartbeat to be backwardly compatible with upgrades to existing cluster #9700
* Externalize Lookup VIndexes properly when not stopping after copy #9771
* Add TEXT field comparison support to evalengine (for VDiff) #9790
* Add CHAR field comparison support to evalengine (for VDiff) #9800
* Run WithDDL queries on tablet startup #9817
* Check for nil vschema.Table in StreamMigrator #9828
* Check if context is cancelled in applyEvents() loop to avoid deadlocks #9833
* Change logic to reach desired _vt schema on vreplication startup #9860
* VPlayer: use stored/binlogged ENUM index value in WHERE clauses #9868
* vstreamer: flatten savepoint events #9892
* vstreamer: do not forward savepoint events #9907
* VReplication: use strict sql_mode for Online DDL #9963
* VStreamer: recompute table plan if a new table is encountered for the same id #9978
* Increase max vrepl workflow source definition size from 64KiB to 16MiB #10018
* VReplication: maintain original column case for pkCols #10033
* Avoid deadlocks related to 0 receiver behavior #10132
* SwitchTraffic should switch everything when no tablet_types provided #10434
* Cherry pick of VStream API: Fix vtgate memory leaks when context gets cancelled #10576
#### VTAdmin
* [vtadmin] Update vtctld dialer to validate connectivity  #9915
* [vtadmin-web] Fix default redirect on Tablet view #10010
* [VTAdmin] Update dynamic clusters to accept no clusters #10016
* [vtadmin] threadsafe dynamic clusters #10044
* [vtadmin-web] Add `PORT=14201` to 'npm run local' to avoid CORS errors in dev mode #10176
#### VTorc
* BugFix: Resolve Recoveries at the end #10286
* Fix panic in VTOrc #10519
#### vtexplain
* Case-insensitive planner-version flag in VTExplain #10086
* fix: check that all keyspaces loaded successfully before using them #10396
### CI/Build
#### Build/CI
* Remove deprecated upgrade-downgrade CI build #9692
* Hack to allow `bootstrap.sh` to run cleanly on fresh M1 #9723
* Fix `make embed_config` target #9724
* Increase async IO threads count in CI #9744
* Fix upgrade-downgrade CI workflows running even when Skip label specified #9784
* Fix flakiness in TestShardedKeyspace #9785
* [grpcvtctldclient] Fix vtctld server initialization in grpcvtctldclient/client_test.go #9906
* Address flakiness of vreplication tests #9935
* Properly generate vtgate_tablet_healthcheck_cache workflow #9950
* Address flakiness of some cluster e2e tests #9967
* Track full git sha1 string instead of the short version #9970
* Set VTDATAROOT to /tmp/ folder in CI #9980
* Flaky TestReparentDuringWorkerCopy fix #9986
* Compatibility Fixes for 1.18 #9992
* Move Shard 12 and 18 CI tests to self hosted runners #10025
* Skip CI workflows based on the list of modified files #10031
* Address flakiness of vreplication v2 workflow test #10051
* makefile: bump vtprotobuf #10078
* Use git command instead of GitHub API to list modified files #10106
* Move tests from newfeaturetest to ers package after the release #10121
* Remove vtadmin-down.sh from teardown.sh #10136
* CODEOWNERS: specify messaging related folders #10137
* bootstrap changelog #10147
* Automatically comment Pull Request with the review checklist #10155
* Applied some format check on the comment_pull_request workflow #10179
* Address flakiness of vtgate_vindex.prefixfanout tests #10216
* Flakes: Add retries to ZooKeeper unit test around startup #10228
* CI: fix flaky test via extended timeout in vrepl_stress_suite #10231
* Revert to using XtraBackup 2.4-latest now that 2.4.26 is out #10241
* Check for vtadmin JS types drift in CI #10256
* Fix `tabletmanager_throttler` flakyness: increase wait time #10271
* CI: Add query retries to vttablet process #10275
* remove references to MySQL and Percona 5.6 #10295
* Aggregate static and other quick CI check actions to reduce #workflows #10431
* fixing vrepl_stress_suite flakyness #10441
* [14.0] Take into account `github.ref` when doing upgrade-downgrade tests #10522
* add vtadmin web files to all lite images #10581
#### Cluster management
* Add GuptaManan100 as a codeowner for VTOrc #9846
* Flakiness fix for Upgrade Downgrade Backup Manual test #9957
* fix endtoend onlineddl flakyness caused by --heartbeat_on_demand_duration #10236
#### Documentation
* Add single self-hosted runner information #9954
#### General
* docker: cleanup of release.sh #9867
* ci: upgrade to Go 1.18 #9893
* 1.18: rename interface{} to any #9914
* Add Matt Lord to CODEOWNERS #9920
* Upgrade main to `go1.18.1` #10101
* Upgrade to `go1.18.3` #10447
#### Governance
* Add @notfelineit to CODEOWNERS #9613
* Update the comment for review checklist with an item for CI workflows #10471
#### Query Serving
* More testing for vreplication unique key violation protection #9988
* unit test: sort map before converting to string #10119
* Fix to --heartbeat_on_demand_duration race condition #10247
* OnlineDDL: adding an endtoend test to validate partial shard REVERT #10338
#### TabletManager
* Alternate fix for flaky TestReparentDuringWorkerCopy #10006
* Fix flakiness in TestHealthCheck #10008
#### VTAdmin
* Add CI workflow for vtadmin-web unit tests #9739
#### VTorc
* Also run VTOrc tests against MySQL 8.0 #9974
### Dependabot
#### Examples
* build(deps): bump gopkg.in/yaml.v2 from 2.2.5 to 2.2.8 in /examples/are-you-alive #9429
#### web UI
* build(deps): bump is-my-json-valid from 2.12.4 to 2.20.6 in /web/vtctld2 #9602
### Documentation
#### Build/CI
* Tweak backport checklist item wording in pull_request_template.md #10083
#### CLI
* [vtctldclient] Update CLI docs for usages, flags, and aliases #10502
#### Documentation
* Update release notes with web/vtctld2 deprecation announcement and flag. #9735
* Update v14 summary with the latest change to `JoinCondition` rewriting #9786
* Add help text for LegacyVtctlCommand and VtctldCommand #9837
* Add durability policy changes to the summary #10392
#### Examples
* Fix commands in readme files to use double dashes for arguments #10389
#### Governance
* update maintainers #9825
* Add @ajm188 to codeowners for CLI-related packages #9847
* Add systay as owner of vtexplain #10403
* docs: added to the release notes #10468
#### Query Serving
* Release notes on query serving #10548
#### VTAdmin
* [vtadmin] Document known issue with node versions 17+ #10483
#### VTorc
* Improve explanation of MySQLTopologyUser and MySQLReplicaUser so that their difference is clearer #9743
### Enhancement
#### Build/CI
* Check for proto drift in the CI #9644
* make generate_ci_workflows: mysql80 workflows for selected tests #9740
* Improve upgrade downgrade workflows' strategy #9745
* Add vitess/lite ubi8 images for mysql80, including for arm64 #9830
* Unit test to confirm that the mysqlctl rice box is current #10182
* Hints to git cli and github diff to hide diffs in proto generated files by default #10306
#### CLI
* [vtctldclient] Add entrypoint for GetVersion rpc #9994
#### Cluster management
* Change semi-sync information to use the parameter passed and deprecate enable_semi_sync #9725
* Filter candidates which cannot make forward progress after ERS #9765
* [vtctldserver] Add locking checks to Delete{Keyspace,Shard} #9777
* [topo] ShardReplicationFix typed responses #9876
* Improve network partitions handling #9905
* tablet lag throttler: small API improvements #10045
* OnlineDDL: skip GetSchema where possible #10107
* Refresh ephemeral information before cluster operations in VTOrc #10115
* On demand heartbeats via `--heartbeat_on_demand_duration`, used by the tablet throttler #10198
* Adds DurabilityPolicy to the KeyspaceInfo in the topo and the associated RPCs #10221
* Refactor Durability Policy implementation and usage to read the durability policy from the keyspace #10375
* Fail if durability policy provided for snapshot keyspaces #10395
* Use Durability Policy in the topo server in VTOrc and deprecate Durability config #10423
* Adds RPCs to vttablet that vtorc requires #10464
#### Documentation
* Add Matt Lord to maintainers #10072
#### Evalengine
* evalengine: Use field information when doing type calculation on columns #9807
#### Examples
* Add consul-topo for local example #9806
* Add vtadmin to local example by default #10430
* Halve request and double limit for RAM in the example cluster #10450
#### General
* Binlog event parsing: better analysis ; support for semi-sync #9596
* Support views in Online DDL #9606
* [VEP-4, phase 1] Flag Deprecation Warnings #9733
* Table lifecycle: support fast DROP TABLE introduced in MySQL 8.0.23 #9778
* Fix rendering type in issue templates #9836
* Flavor capabilities: MySQL GR and more #10451
#### Query Serving
* Harmonize error codes for pool timeouts expiring across all pools #9483
* Store innodbRowsRead in a Counter instead of a Gauge #9609
* Improve performance of information schema query. #9632
* OnlineDDL: reject duplicate UUID with different migration_context #9637
* Reduce the number of reserved connections when setting system variables #9641
* OnlineDDL: 'vitess' strategy, synonym for 'online' #9642
* Gen4: Rework how aggregation is planned #9643
* evalengine: More functions! #9673
* Support for ALTER TABLE ... PARTITION BY ... #9683
* OnlineDDL: force (once) vreplication's WithDDL to run before running a vitess migration #9702
* Push projection to union #9703
* Make gen4 the default planner for release 14 #9710
* Cleanup of parsing of Partitions and additional parsing support #9712
* Introducing schemadiff, a declarative diff for table/view CREATE statements #9719
* Online DDL vitess migration's cut-over: query buffering #9755
* Online DDL declarative migrations now use schemadiff, removing tengo #9772
* Online DDL: ready_to_complete hint column #9813
* Add parsing support for prepare statements #9818
* mysql: add a flag to handle EnableQueryInfo #9820
* Online DDL: vitess migrations cut-over to have zero race conditions #9832
* Add parsing support for Trim grammar function #9834
* Add parsing support for LATERAL keyword #9843
* Add parsing support for JSON Utility Functions #9850
* Remove the `Gateway` interface #9852
* feat: make the join columns output easier to read #9863
* Gen4: Add UPDATE planning #9871
* Add parsing support for JSON_TABLE() #9879
* Add parsing support for JSON value creators functions #9880
* fix: allow multiple distinct columns on single shards #9940
* Add parsing support for Json schema validation functions #9971
* Add parsing support for JSON Search Functions #9990
* deps: upgrade grpc and protobuf #10024
* feat: allow more complex expressions to be used in outer queries #10034
* schemadiff: load, normalize, validate and compare schemas #10048
* Add parsing support for JSON value modification and attribute functions #10062
* Always reset reserved shard session on CODE_UNAVAILABLE #10063
* Adds parsing for NOW in DEFAULT clause #10085
* Gen4 feature: optimize OR queries to IN #10090
* Relax singleton-context constraint for pending non-singleton-context reverts #10114
* Add parsing support for Partition Definitions #10127
* feat gen4: allow pushing aggregations inside derived tables #10128
* schemadiff: support for RangeRotationStrategy diff hint #10131
* Avoid enabling reserved connection on show query #10138
* Add support for non aggregated columns in `OrderedAggregate` #10139
* schemadiff: CanonicalStatementString(), utilize sqlparser.CanonicalString() #10142
* feat: make gen4 plan subsharding queries better #10151
* schemadiff: EntityDiff has a 'Entities() (from Entity, to Entity)' function #10161
* Introducing a SQL syntax to control and view Online DDL throttling #10162
* feat: optimize EXISTS queries through AST rewriting #10174
* schemadiff: logically Apply() a diff to a CreateTable, CreateView or to a Schema #10183
* Using LCS to achieve shortest column reordering list #10184
* enhancement: handle advisory locks  #10186
* schemadiff: normalize() table. Set names to all keys #10188
* schemadiff: validate() table structure at the end of apply() #10189
* schemadiff: schema is immutable through Apply() #10190
* VReplication: use PK equivalent (PKE) if no PK found #10192
* Add parsing support for Window Functions #10199
* schemadiff: normalize table options case #10200
* Allow updating unowned lookup vindex columns. #10207
* partial dml execution logic ehancement #10223
* gen4: JOIN USING planning #10226
* schemadiff: partitions validations #10229
* Add parsing support for Subpartition Definitions #10232
* schemadiff: analyze and apply ADD PARTITION and DROP PARTITION statements #10234
* Add convenience functions for working with schemadiff objects #10238
* Add parsing support for expressions in index definitions #10240
* Update the parsing support for JSON_VALUE #10242
* schemadiff: support functional indexes #10244
* Fix formatting of character set annotation on string literals #10245
* schemadiff: use ParseStrictDDL whenever expecting a DDL statement #10246
* Fix parsing of encryption attribute for create database #10249
* Add support for parsing MATCH in a foreign key definition #10250
* Add parsing support for additional index and table options #10251
* Handle additional DDL in column definitions #10253
* evalengine: CASE expressions #10262
* vt/sqlparser: support SRID #10267
* Improve dealing with check constraints #10268
* Add parsing and schemadiff support for ALTER TABLE ALTER CHECK #10269
* Add unary utf8mb3 charset marker #10274
* Allow updating the subsequent diff #10289
* Improve the formatting for partitioning #10291
* Add additional normalization rules for schemadiff #10296
* schemadiff: RangeRotationDistinctStatements partitioning hint #10309
* Remove internal savepoint post query execution #10311
* feat: use column offsets when HAVING predicate has aggregation #10312
* Improve parser and schemadiff support for ALTER TABLE #10313
* Online DDL: fast (and correct) RANGE PARTITION operations, 1st iteration #10315
* Add support for additional encoding configurations #10321
* Add parsing support for Regular Expressions #10322
* schemadiff: validate columns referenced by generated columns #10328
* schemadiff: case insensitive col/index/partition name comparison #10330
* [gen4] More aggregation support #10332
* Add better int normalization #10340
* Additional schemadiff improvements for indexes #10345
* Structured schemadiff errors #10356
* Add validation of check constraints #10357
* schemadiff: validate columns referenced by FOREIGN KEY #10359
* Aggregation on top of LIMIT #10361
* Add support for parsing ARRAY in cast / convert operations #10362
* Online DDL: support for CHECK CONSTRAINTS #10369
* schemadiff: constraint name diff hints strategies #10380
* Replace 'CREATE TABLE LIKE' with programmatic copy #10381
* Streaming implementation of Projection engine primitive #10384
* feat: do not use lookup vindexes for IS NULL predicates #10388
* Online DDL: support INSTANT DDL special plan, flag protected (flag undocumented for now) #10402
* feat: make v3 not plan IS NULL queries against a lookup vindex #10412
* Online DDL: support Instant DDL for change of column's DEFAULT #10414
* Parsing Support for XML functions #10438
* Formalize MySQL capabilities by flavor/version #10445
* fix: change planner_version to planner-version everywhere #10453
* [14.0] Schema tracking acl error logging #10591
* [Backport 14.0] enable schema tracking by default  #10595
#### TabletManager
* messaging: support vt_message_cols to limit cols #9670
* Table lifecycle: support views #9776
#### VReplication
* OnlineDDL: delete _vt.vreplication entry as part of CLEANUP operation #9638
* Tablet server/VReplication: use information_schema.columns to get the list of columns instead of "select * from table"  #9794
* Add `--drop_constraints` to MoveTables #9904
* Support throttling vstreamer copy table work on source tablets #9923
* Improve escape handling for vindex materializer #9929
* VReplication: fail on unrecoverable errors #9973
* Prefer using REPLICA tablets when selecting vreplication sources #10040
* Improve Tablet Refresh Behavior in VReplication Traffic Switch Handling #10058
* VReplication: use db_filtered user for vstreams #10080
* MoveTables: adjust datetimes when importing from non-UTC sources into UTC targets #10102
* VStream API: allow cells to be specified for picking source tablets to stream from #10294
* Vdiff2: initial release #10382
* release-14.0 backport: Fail VReplication workflows on errors that persist and unrecoverable errors #10573
* Backport: VDiff2: Support Resuming VDiffs #10589
#### VTAdmin
* [vtadmin-web] Upgrade to tailwindcss v3 #9648
* [vtadmin-web] Add REACT_APP_READONLY_MODE flag #9731
* [VTAdmin] URL decode cookies #9866
* [vtadmin] custom discovery resolver #9977
* [vtadmin-web] Add /vschema route to Keyspace detail view + QueryErrorPlaceholder and QueryLoadingPlaceholder components #10021
* [vtadmin] grpc healthserver + channelz #10038
* [Dynamic Clusters] No Duplicates #10043
* [vtadmin] full dynamic cluster config #10071
* [VTAdmin] Replace changed clusters when using Dynamic Discovery #10129
* [vtadmin] non-blocking resolver #10205
* [vtadmin] Add "Replication Status" view to Tablet page #10266
* [vtadmin] Add support for 'allow_primary' when deleting a tablet #10304
* [vtadmin] Add CreateKeyspace bindings + form UI #10316
* [vtadmin] tablet topopools + other cleanup #10325
* [vtadmin] Add keyspace/advanced view + ReloadSchemaKeyspace control #10329
* [vtadmin] Add loading placeholder for entity table views + update default redirect to /schemas #10331
* [vtadmin] cluster Closer implementation + dynamic Cluster bugfix #10355
#### VTCombo
* Add flag to vtcombo to load a json-encoded test topology file #9633
#### VTorc
* Revocation phase using durability policies #9659
* VTOrc: Refactor and reload of ephemeral information for remaining recovery functions #10150
* vtorc: more auditing in GracefulPrimaryTakeover #10285
#### vttestserver
* vtcombo: Add the ability to specify routing rules in test topology #9695
* vttestserver: add enable_system_settings option #9779
#### web UI
* [vtctld] Add 'enable_vtctld_ui' flag to vtctld to explicitly enable (or disable) the vtctld2 UI #9614
### Feature Request
#### Cluster management
* [vtctldserver] Migrate `Backup` RPC to vtctldserver/client #9798
* [vtctldclient] Add entrypoint for `ApplySchema` #9829
* [vtctldserver] Migrate `BackupShard` RPC #9857
* [vtctldserver] Migrate `RemoveBackup` RPC #9865
* [vtctldserver] Migrate `RestoreFromBackup` RPC #9873
* [vtctldserver] Migrate `ShardReplicationFix` RPC #9878
* [vtctldserver] Migrate `SourceShard{Add,Delete}` RPCs #9886
* [vtctldserver] ExecuteFetchAs{App,DBA} #9890
* [vtctldserver] Migrate `ShardReplication{Add,Remove}` RPCs #9899
* [vtctldserver] Migrate GetPermissions PR#9903 #10013
#### Observability
* [vtadmin] Add prom metrics endpoint support #10334
#### Query Serving
* mysql: allow parsing the 'info' field of MySQL packets #9651
* evalengine: Conversions #9698
* Add support for insert using select #9748
* Add parsing support for Multi INDEX HINTS #9811
* sqlparser: introduce custom formatting options #10068
#### VTAdmin
* [VTAdmin] Tablet Actions (vtctld2 Parity) #9601
* [VTAdmin][vtctldserver] Keyspace Functions #9667
* Allow --no-rbac flag that allows users to not pass rbac config #9972
* [vtadmin] grpc dynamic clusters #10050
* [vtadmin] schema cache #10120
* [vtadmin] cell apis #10227
* [vtadmin] reload schema #10300
* [vtadmin] reparenting actions #10351
* add vtadmin docker image #10543
### Internal Cleanup
#### Build/CI
* Update the issue templates #9605
* Remove usage of additional uuid package #10307
#### CLI
* [vtctl] Deprecate query commands #9934
* [vtctl] Deprecate Throttler RPCs #9962
* [vtadmin] Remove deprecated dynamic cluster flag #10067
* Update vtctl help output to use double dashes for long flag names #10405
#### Cluster management
* Remove the DurabilityParams argument to the durability interface #9718
* Internal cleanup for choosing primary in ERS #9775
* Move deprecation warning to just before we actually invoke the legacy command path #9787
* Online DDL: cleanup old vtctld+topo flow #9816
* Rename ApplySchemaRequest.request_context=>migration_context #9824
* rename master_alias to primary_alias in reparent_journal table #10098
* Use the `HealthCheckImpl` in `vtctld` instead of the `LegacyHealthCheck` #10254
* Cleanup: removed legacy, unused code (online DDL paths in topo) #10379
* Addition of more logging to ERS and tests to PRS #10425
#### Documentation
* all: fix some typos #9815
#### Evalengine
* evalengine: more comprehensive type coverage #9804
#### Examples
* Examples: fix warnings #9875
* Migrate remaining vtctldclient commands in local example #9894
* [vtadmin] Improve vtadmin-{up,down}.sh scripts for the local example #10081
#### General
* Support sanitizing potentially sensitive information in log messages #9550
* Update maintainer email address for Andrew Mason #9709
* Remove unmaintained vagrant files #9768
* [VEP-4, phase 1] Flag cleanup in `go/cmd` #9896
* Add VS Code's __debug_bin to .gitignore #10305
* fix: lint error #10440
#### Query Serving
* Delete discovery gateway #9500
* Remove queryserver-config-terse-errors impact on log messages #9634
* evalengine: use sqltypes #9646
* Add a comment on how to use set for a counter #9664
* evalengine: Value cleanups #9773
* Split aggregation in scalar and ordered aggregation #9810
* refactor: move all remaining queries to plan builder #9982
* Minor fix requested by linter #9989
* collations/hack: cleanup dependencies #10039
* Reduce shift-reduce conflicts in the grammar by using Precedence #10053
* remove set_var session field from proto #10135
* refactor: removed last reserved field from vtgate session #10140
* refactor: move the Subshard opcode #10195
* fix: filter was not using the vcursor buffering execution #10391
#### TabletManager
* Adds comments to UnserveCommon to find potential Deadlock #9942
* Fix Flaky TestTopoCustomRule #10002
* delete RPCs that were deprecated in 13.0 #10099
#### VReplication
* Rename drop_constraints to drop_foreign_keys on MoveTables #10017
* remove changeMasterToPrimary, we no longer need it #10100
* Ensure all legacy sharding commands are clearly deprecated #10281
#### VTAdmin
* [vtadmin-web] Update postcss to v8.4.6 #9612
* [VTAdmin] Simplify JSON Discovery #9618
* [vtadmin-web] Upgrade stylelint + related dependencies to latest #9654
* [vtadmin-web] Update msw to latest (v0.36.8) #9737
* chore(deps): Bump node-forge from 1.2.0 to 1.3.0 in /web/vtadmin #9953
* [vtadmin] threadsafe vtsql #10000
* [vtsql] Remove vtsql.PingContext check from dial calls #10037
* [vtadmin-web] Update to react-scripts@5.0.1 #10082
* [vtadmin] Refactor api.getSchemas => cluster.GetSchemas #10108
* [vtadmin] Remove Tablets option for GetSchema(s) #10111
* [vtadmin-web] Remove placeholder links from NavRail #10180
* [vtadmin-web] Remove unused icon SVGs #10206
* [vtadmin] Tidy up /settings (née /debug) page #10218
* [vtadmin] Remove explicit Dial-ing for vtctldclient proxy #10233
* [vtadmin] Removing explicit Dial-ing for vtsql proxy #10237
* [vtadmin] `GetCell{Infos,Aliases}`: Add missing topoReadPool management #10273
* [vtadmin] proper alias types for tablet RPCs #10290
* [vtadmin] Rehome DangerAction component in a common folder #10293
* [vtadmin] Fix warning 'div cannot appear as a descendant of p' #10302
* Mark web/vtadmin/src/proto/** as generated files in .gitattributes #10303
* [vtadmin] Further generalize DangerAction component #10319
* [vtadmin] Further generalize DangerAction into ActionPanel #10323
* [vtadmin] Remove source-map-explorer dependency #10339
* [vtadmin] Refactor getTabletForAction #10341
* [vtadmin] Update ejs, async, and webpack-dev-server dev dependencies #10342
* [vtadmin] Remove mutex, now that vtexplain has no global state #10387
* [vtadmin] Rename reparent-related RPCs, endpoints, methods, RBAC actions #10404
* [vtadmin] Rename ERS/PRS pools+flags properly #10460
#### VTorc
* Remove logging in GetDurabilityPolicy #10516
#### vtctl
* Fixes #8277: vtctld logs leak pii http headers #9669
#### vtexplain
* Vtexplain no globals #10374
* replace planner-version with planner_version in vtexplain #10420
#### web UI
* [VTAdmin Web] Make Protos #10042
### Performance
#### Build/CI
* Replace OpenSSL usage for test certificates with Go crypto #9999
#### Examples
* region_example: sleep 5s before switching reads to allow copy to finish #9937
#### General
* sqlescape: Improve performance of EscapeID #9840
#### Query Serving
* Improve performance of hex type decoding #9625
* Micro benchmarks for setting system variables #9681
* evalengine: Optimize NullSafeComparison #9728
* evalengine: Decimals v2 Electric Boogaloo #9749
* sqlparser: cache comment parsing #10014
#### Schema Tracker
* Improve performance of the `DetectSchemaChange` query #10416
### Release
#### Build/CI
* Updating 13.0.0 release notes on main #9759
* Rework how the `release notes` labels are handled by the CI #10508
* Rework the generation of the release notes #10510
* Post release 14.0.0 updates #10600
#### General
* Moving dev to 14.0.0-SNAPSHOT #9611
* Addition of release notes for v13.0.1 #10097
* add caller_id to 13.0.0 release notes #10124
* Release notes documents for `v12.0.4` on `main` #10350
* [14.0] Update and prepare the `v14.0.0` summary #10569
* Revert snapshot incrementation in release-14.0 #10605
### Testing
#### Build/CI
* More aggressive tests for vitess migration cut-over #9956
* Skip end-to-end test if binary is below a certain major version #9975
* Deflake connkiller tests #10001
* Temp: Pin XtraBackup version used at 2.4.24 for 5.7 tests #10194
* fix tablegc flaky test #10230
#### CLI
* [VEP-4, part 1] Flag cleanup in `go/mysql` and some endtoend tests #9910
* [VEP4, phase 1] Update flags in `endtoend/recovery` #9911
* [VEP4, phase 1] Update flags in `endtoend/schemadiff` #9912
* [VEP4, phase 1] Update flags and imports for `endtoend/{sharded,sharding}` #9913
* [VEP 4, phase 1] Cleanup flags in `endtoend/vault` (and some comments in `endtoend/tabletmanager`) #9916
* [VEP4, phase 1] fix flags in tests `endtoend/versionupgrade` #9917
* [VEP4, phase 1] [endtoend/vreplication] update flags in tests #9918
* [VEP4, phase 1] [endtoend/vtcombo] fix flags in tests #9919
* [VEP4, phase1] Cleanup flags in `endtoend/vtgate` and `endtoend/vtorc` #9930
* [VEP4, phase1] fix flags in tests in `endtoend/worker` and strings in `go/trace` #9931
#### Cluster management
* Replace InitShardPrimary usage in tests by PlannedReparentShard #9680
#### General
* Add unit test for the demo app #9766
* test: use `T.TempDir` to create temporary test directory #10433
#### Query Serving
* Add e2e test for the TabletManager GetSchema RPC #9658
* vtexplain: added multicol test to show subshard routing #9697
* Compare Vitess against MySQL in E2E tests #9965
* endtoend: improve assertionsin message_test #10046
* Flaky test fix: TestFoundRows #10235
* Make the tests explicit for the SET variables for MySQL 5.7 #10263
* fix: vtexplain commit order flaky #10264
* test refactor: extract const literal to file #10358
* test: refactor and enabled vschema ddl test #10437
* Divide query serving upgrade downgrade tests in two groups to limit disk usage per workflow #10452
#### VReplication
* deflake TestCanSwitch by switching from exact string match to loose regex #9926
#### VTAdmin
* [vtadmin] Ensure we close the proxy to prevent red-herring 'error' messages #9964
* [vtadmin|tests] Add mutex to prevent data races in tests using cluster/resolver #10041
* [vtadmin] tablet topotools tests #10333
* [vtadmin] Fix data race in vtadmin #10376
* [vtadmin] Add infrastructure for generating authz tests for vtadmin #10397
* [vtadmin] authz tests part1 #10415
* [vtadmin] Refactor authztestgen template to allow multiple test clusters #10428
* [vtadmin] authz tests - getters part1 #10432
* [vtadmin] authz tests - getters part2 #10444
* [vtadmin] authz tests -  getters part3 #10449
* [vtadmin] authz tests - tablet actions #10457
* [backport] [vtadmin] Add remaining authz tests #10487
* [vtadmin/tests] Serialize Schema test cases to avoid cache backfill races #10538
* [vtadmin] fix flaky GetSchemas test cases #10555
#### VTorc
* Remove the use_super_read_only flag in VTOrc tests #9672
