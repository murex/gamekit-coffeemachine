#!/usr/bin/env bash

base_dir="$(cd "$(dirname -- "$0")" && pwd)"

path_to_implementation="$1"
if [ -z "${path_to_implementation}" ]; then
  echo "Usage: $0 <path to implementation>"
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
# Assumption: the last directory in the path is the language name
language=$(basename "${path_to_implementation}")
echo "running progress tests against ${language} implementation of Coffee Machine"

export LANG_IMPL_PATH="${path_to_implementation}"

mkdir -p _test_results

gotestsum_exe="${base_dir}/bin/gotestsum${extension}"
test2json_exe="${base_dir}/bin/test2json${extension}"
progress_runner_exe="${base_dir}/bin/progress-runner${extension}"

package_name=progress-runner-"${language}"
junitfile=_test_results/progress_runner-"${language}".xml
if [ "${trace_mode}" == "-vv" ]; then
  # Prints test names and failed tests output
  ${gotestsum_exe} \
    --format testdox \
    --junitfile "${junitfile}" \
    --raw-command \
    -- "${test2json_exe}" -t -p "${package_name}" "${progress_runner_exe}" -test.v=test2json
elif [ "${trace_mode}" == "-v" ]; then
  # Prints test names but not failed tests output
  ${gotestsum_exe} \
    --format testdox \
    --hide-summary=all \
    --junitfile "${junitfile}" \
    --raw-command \
    -- "${test2json_exe}" -t -p "${package_name}" "${progress_runner_exe}" -test.v=test2json
else
  # Prints dots only
  ${gotestsum_exe} \
    --format dots \
    --hide-summary=all \
    --junitfile "${junitfile}" \
    --raw-command \
    -- "${test2json_exe}" -t -p "${package_name}" "${progress_runner_exe}" -test.v=test2json
fi
