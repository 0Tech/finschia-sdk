set(SIMD_NUM_INSTANCES 7 CACHE STRING "The number of simd instances")
set(SIMD_MODE test CACHE STRING "simd mode, options are: test")
set(SIMD_NET_BEGIN 192.168.0.2/16 CACHE STRING "the start address of simd network")

set(SIMD_TIMEOUT_COMMIT 1s CACHE STRING "timeout_commit of simd")
set(SIMD_TIMEOUT_PRECOMMIT 1s CACHE STRING "timeout_precommit of simd")
set(SIMD_TIMEOUT_PREVOTE 1s CACHE STRING "timeout_prevote of simd")
set(SIMD_TIMEOUT_PROPOSE 1s CACHE STRING "timeout_propose of simd")

configure_file(generate-simd.py.in generate-simd.py
  @ONLY)
configure_file(simd.env.in simd.env
  @ONLY)

add_custom_command(
  OUTPUT simd.yml
  DEPENDS generate-simd.py
  COMMAND ./generate-simd.py simd-${SIMD_MODE} ${SIMD_NUM_INSTANCES} --network ${SIMD_NET_BEGIN} -o simd.yml
  VERBATIM)
get_target_property(INSTALL_SIMD_DIR install_simd BINARY_DIR)
add_custom_command(
  DEPENDS simd.yml install_simd ${INSTALL_SIMD_DIR}/install-simd.touch
  OUTPUT deploy-simd.touch
  COMMAND sh -c "COMPOSE_PARALLEL_LIMIT=${SIMD_NUM_INSTANCES} docker-compose -f simd.yml -p simd-${SIMD_MODE} up -d --remove-orphans; touch deploy-simd.touch"
  COMMENT "Launching Docker instances for simd"
  VERBATIM)
add_custom_target(deploy_simd
  DEPENDS deploy-simd.touch)
add_custom_target(undeploy_simd
  COMMAND sh -c "[ -f deploy-simd.touch ] && docker-compose -f simd.yml -p simd-${SIMD_MODE} down -v; rm -f deploy-simd.touch"
  VERBATIM)
