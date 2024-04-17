# How to reference the tile set

## Grass.png

The grass png file is supposed to be referenced in 15x15 pixel areas.
0,0,15,15 meaning 0px from left, 0px from top, 15px to the left and 15 px down.

the map is
175 wide and 111 high. Not all tiles in the sprite would have value - so you should test and make sure the source hits well.

edge left-top : 0,0
edge top: 16, 0
edge top-right: 32, 0
edge left-top-right: 48, 0
edge left-top-right(corner bottom-right): 64, 0
top + corner bottom-right: 80, 0
top + corner bottom-left: 96, 0
edge top-right + corner bottom-left: 112, 0
edge top + corner bottom-left, corner bottom-right: 128, 0
144 corner top-left, bottom-right
160 (empty)

We will use these 9 for the tutorial:

x x x | x x x | x x x | 0,0 | 16,0 | 32,0 |
x o o | o o o | o o x |  
x o o | o o o | o o x |

x o o | o o o | o o x | 0,16 | 16,16 | 32,16
x o o | o o o | o o x |
x o o | o o o | o o x |

x o o | o o o | o o x | 0,32 | 16,32 | 32,32
x o o | o o o | o o x |
x x x | x x x | x x x |
