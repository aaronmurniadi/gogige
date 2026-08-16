|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

**: An integer in the range iof {∅, 2} declaring the number of consecutive pixels in vertical direction that are outputted simultaneously from each zone. This field is omitted when all pixels are in the same line.

**: A letter in range of {∅, E} declaring the location of the pixels extractors in the vertical direction. The value E indicates that pixel extractors are at both top and bottom lines. This field is omitted when all pixel extractors are in the top line.

### 29.2.1.3 Tap Geometrical Properties

The following tables provide description of all the tap geometry configurations. For every configuration the first and last pixel belonging to that tap, as well as the pixel increment corresponding to the given tap is listed.

This table enumerates the standard tap geometries. The table is sorted by increasing number of taps. It displays the values of the 6 geometrical properties for each tap.

Table 29-1 Tap geometrical properties – One, two and three taps

|  Geometry name |   | Tap | Tap geometrical properties  |   |   |   |   |   |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
|  Area-scan | Line-scan |   | X Start | X End | Step X | Y Start | Y End | Step Y  |
|  1X-1Y | 1X | Tap1 | 1 | W | 1 | 1 | H | 1  |
|  1X2-1Y | 1X2 | Tap1 | 1 | W-1 | 2 | 1 | H | 1  |
|   |   |  Tap2 | 2 | W | 2 | 1 | H | 1  |
|  2X-1Y | 2X | Tap1 | 1 | W/2 | 1 | 1 | H | 1  |
|   |   |  Tap2 | W/2+1 | W | 1 | 1 | H | 1  |
|  2XE-1Y | 2XE | Tap1 | 1 | W/2 | 1 | 1 | H | 1  |
|   |   |  Tap2 | W | W/2+1 | -1 | 1 | H | 1  |
|  2XM-1Y | 2XM | Tap1 | W/2 | 1 | -1 | 1 | H | 1  |
|   |   |  Tap2 | W/2+1 | W | 1 | 1 | H | 1  |
|  1X-1Y2 |  | Tap1 | 1 | W | 1 | 1 | H-1 | 2  |
|   |   |  Tap2 | 1 | W | 1 | 2 | H | 2  |
|  1X-2YE |  | Tap1 | 1 | W | 1 | 1 | H/2 | 1  |
|   |   |  Tap2 | 1 | W | 1 | H | H/2+1 | -1  |
|  1X3-1Y | 1X3 | Tap1 | 1 | W-2 | 3 | 1 | H | 1  |
|   |   |  Tap2 | 2 | W-1 | 3 | 1 | H | 1  |
|   |   |  Tap3 | 3 | W | 3 | 1 | H | 1  |
|  3X-1Y | 3X | Tap1 | 1 | W/3 | 1 | 1 | H | 1  |
|   |   |  Tap2 | W/3+1 | 2W/3 | 1 | 1 | H | 1  |
|   |   |  Tap3 | 2W/3+1 | W | 1 | 1 | H | 1  |