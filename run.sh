#!/usr/bin/env bash

base_dir="$(cd "$(dirname -- "$0")" && pwd)"

language="$1"
if [ -z "$language" ]; then
  echo "Usage: $0 <language>"
  exit 1
fi

trace_mode="$2"

case  $(uname -s) in
  Darwin) extension="";;
  Linux) extension="";;
  MINGW64_NT-*) extension=".exe";;
  *) echo "os $(uname -s) not supported." && exit 1;;
esac

#set -eu

# 1) Build progress test executables (one per iteration)

echo "building progress test executables"
if [ -z "${trace_mode}" ]; then
  make -C "${base_dir}" --silent build
else
  make -C "${base_dir}" build
fi

# 2) Run progress tests through gotestsum
echo "running progress tests against ${language} implementation of Coffee Machine"

export RUN_ON_LANG="${language}"

mkdir -p _test_results

gotestsum_exe="${base_dir}/bin/gotestsum${extension}"
test2json_exe="${base_dir}/bin/test2json${extension}"
progress_tests_exe="${base_dir}/bin/progress-tests${extension}"

package_name=progress-tests-"${language}"
junitfile=_test_results/progress_tests-"${language}".xml
if [ "${trace_mode}" == "-vv" ]; then
  # Prints test names and failed tests output
  ${gotestsum_exe} \
    --format testdox \
    --junitfile "${junitfile}" \
    --raw-command \
    -- "${test2json_exe}" -t -p "${package_name}" "${progress_tests_exe}" -test.v=test2json
elif [ "${trace_mode}" == "-v" ]; then
  # Prints test names but not failed tests output
  ${gotestsum_exe} \
    --format testdox \
    --hide-summary=all \
    --junitfile "${junitfile}" \
    --raw-command \
    -- "${test2json_exe}" -t -p "${package_name}" "${progress_tests_exe}" -test.v=test2json
else
  # Prints dots only
  ${gotestsum_exe} \
    --format dots \
    --hide-summary=all \
    --junitfile "${junitfile}" \
    --raw-command \
    -- "${test2json_exe}" -t -p "${package_name}" "${progress_tests_exe}" -test.v=test2json
fi
