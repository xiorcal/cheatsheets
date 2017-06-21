# Bash

## Language convention

## tooling

### CLI tools

#### beautysh ([github](https://github.com/bemeurer/beautysh))

This program takes upon itself the hard task of beautifying Bash scripts (yeesh).

install with `pip install beautysh`

#### shellcheck ([github](https://github.com/koalaman/shellcheck))

A shell script static analysis tool

install with `cabal update ; cabal install ShellCheck`

### relevant atom packages

- atom-beautify
- linter
- linter-shellcheck

## tips

### run processes in parallel

```bash
#!/bin/bash
for c in $commands_to_run
do
    c &
done

FAIL=0

echo "start waiting at $(date +%s)"
for job in $(jobs -p)
do
  echo "$job"
  wait "$job" || let "FAIL+=1"
done

echo "$FAIL"
echo "done wait (TS : $(date +%s))"
```
