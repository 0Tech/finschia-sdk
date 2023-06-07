set(FIXTURE_DAEMON_VERSION_NEW dummy CACHE STRING "new binary version for the upgrade (dummy|local)")

get_target_property(IMAGE_SOURCE_DIR install_image SOURCE_DIR)
get_target_property(DAEMON_BINARY_DIR build_daemon BINARY_DIR)

add_subdirectory(rolling)
add_subdirectory(in-place)
