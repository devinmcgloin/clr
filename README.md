# clr

clr is a go library to manage different color spaces, convert between them and
compare colors.

clr currently supports sRGB and Hex types, with conversions to sRGB, Hex, HSV,
HSL, CMYK, XYZ and CIE-Lab for each.

Comparisons are done in CIE-Lab using Delta E*. I may add more comparison
options in the future, but for now Delta E* performs well enough.

There is also a CLI tool to explore this package called `clr`. 

```
$ clr --hex "#0c6624"
&{Code:0c6624}
sRGB: 12 102 36
HSL: 136 78 22
HSV: 136 88 40
CMYK: 88 0 64 60
HEX: 0c6624
Generic Color: Green
Specific Color: Camarone
```

