set(SIMD_NUM_INSTANCES 7 CACHE STRING "The number of simd instances")
set(SIMD_KEY_MNEMONIC "mind flame tobacco sense move hammer drift crime ring globe art gaze cinnamon helmet cruise special produce notable negative wait path scrap recall have" CACHE STRING "mnemonic for keys")
set(SIMD_NET_BEGIN 192.168.0.2/16 CACHE STRING "the start address of simd network")

set(SIMD_CHAIN_ID test CACHE STRING "simd chain id")
set(SIMD_KEYRING_BACKEND test CACHE STRING "simd keyring backend, options are: os, file, test")

set(SIMD_TIMEOUT_COMMIT 1s CACHE STRING "timeout_commit of simd")
set(SIMD_TIMEOUT_PRECOMMIT 1s CACHE STRING "timeout_precommit of simd")
set(SIMD_TIMEOUT_PREVOTE 1s CACHE STRING "timeout_prevote of simd")
set(SIMD_TIMEOUT_PROPOSE 1s CACHE STRING "timeout_propose of simd")

set(SIMD_ENABLE_CONSORTIUM false CACHE STRING "enable consortium")

configure_file(generate-simd.py.in generate-simd.py
  @ONLY)
configure_file(simd.env.in simd.env
  @ONLY)
configure_file(dsimd.in dsimd
  @ONLY)

add_custom_command(
  OUTPUT simd.yml
  DEPENDS generate-simd.py
  COMMAND sh -c "[ -f simd.yml ] && ./dsimd down || true; ./generate-simd.py ${SIMD_NUM_INSTANCES} --network ${SIMD_NET_BEGIN} -o simd.yml"
  COMMENT "Generating Docker compose file for simd"
  VERBATIM)
get_target_property(INSTALL_SIMD_DIR install_simd BINARY_DIR)
add_custom_target(deploy_simd
  DEPENDS simd.yml install_simd ${INSTALL_SIMD_DIR}/install-simd.touch)
