configure_file(shutdown-sentries.sh.in shutdown-sentries.sh
  @ONLY)

add_test(
  NAME shutdown_sentries
  COMMAND timeout 5m sh shutdown-sentries.sh 0)
set_tests_properties(shutdown_sentries
  PROPERTIES FIXTURES_REQUIRED CHAIN)
