cmake_minimum_required(VERSION 3.24)
project(array)

set(CMAKE_CXX_STANDARD 11)

include_directories(.)

add_executable(array ArrayList.hpp main.cpp)