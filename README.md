# percentaverage
A tool which finds all precentages in a file and averages them. Matches `x.y%`, `x,y%`, `-x.y%`.

### Installation
```bash
go install github.com/baalimago/percentaverage@latest
```

## Usage

```bash
echo "20,3% 19.9% -0.2%" | percentaverage && echo
```

```bash

go test ./... -cover | percentaverage -r `# This will output the percentage wihout a '%' sign`
```

## Roadmap
- [x] Parse stdin
- [ ] Hyperdrive the parsing for fun (fun not guaranteed)
- [ ] Parse file from `-f` flag
- [ ] Parse all files from directory using `-d` flag
- [ ] Paralellized parsing
- [ ] Glob parsing using `-g` flag
- [ ] Glob parsing using multi-arguments
