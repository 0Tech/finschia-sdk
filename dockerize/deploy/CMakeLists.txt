configure_file(generate-compose.py.in generate-compose.py
  @ONLY)
configure_file(dsimd.in dsimd
  @ONLY)

add_custom_command(
  OUTPUT deploy.touch
  DEPENDS generate-compose.py dsimd
  COMMAND touch deploy.touch
  COMMENT "Generating files for deploy"
  VERBATIM)
add_custom_target(deploy
  DEPENDS deploy.touch install_image)
