---
name: color-contrast
description: "Calculate WCAG color contrast ratios"
---
# Color Contrast Checker

Calculate WCAG contrast ratios using **bash** + Python (or inline math).

## Quick Check (Python one-liner)
```bash
python3 -c "
def lum(r,g,b):
    cs = [c/255.0 for c in (r,g,b)]
    cs = [(c/12.92 if c<=0.03928 else ((c+0.055)/1.055)**2.4) for c in cs]
    return 0.2126*cs[0]+0.7152*cs[1]+0.0722*cs[2]
def contrast(fg,bg):
    l1,l2 = max(lum(*fg),lum(*bg)),min(lum(*fg),lum(*bg))
    return (l1+0.05)/(l2+0.05)
fg = (FG_R, FG_G, FG_B)
bg = (BG_R, BG_G, BG_B)
r = contrast(fg,bg)
print(f'Ratio: {r:.2f}:1')
print(f'AA Normal (4.5:1): {\"PASS\" if r>=4.5 else \"FAIL\"}')
print(f'AA Large (3:1): {\"PASS\" if r>=3.0 else \"FAIL\"}')
print(f'AAA Normal (7:1): {\"PASS\" if r>=7.0 else \"FAIL\"}')
"
```

## WCAG Requirements
| Level | Normal Text | Large Text |
|-------|-------------|------------|
| AA | 4.5:1 | 3:1 |
| AAA | 7:1 | 4.5:1 |

Large text = 18pt+ regular or 14pt+ bold.

## Tips
- Convert hex to RGB: #FF6600 = (255, 102, 0)
- Test both light and dark mode combinations
- Common fail: light gray text on white background
