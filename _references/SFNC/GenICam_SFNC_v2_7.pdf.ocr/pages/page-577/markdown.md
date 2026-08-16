|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

○ YStep: Difference of Y-coordinates between consecutive pixel rows; Y-step is positive when Y-coordinates is increasing at the end of a line; it is negative otherwise.
○ Allocation of taps to ports. The camera taps are indexed using following conventional sorting rule:
First by increasing values of YStart then by increasing value of XStart. The tap T1 is the sensor tap that exhibits the smallest XStart for the smallest YStart.

### 29.2.1.1 Restrictions

- All zones have the same size.
- Zones do not overlap.
- All zones have the same number of taps.
- All taps are carrying the same amount of pixels.

### 29.2.1.2 Tap Naming Convention

A tap configuration for area-scan camera is designated by:

<TapGeometryX>-<TapGeometryY></TapGeometryX>

A tap configuration for line-scan or TDI-line-scan camera is designated by:

<TapGeometryX></TapGeometryX>

TapGeometryX is designated by <ZoneX>X(<TapX>)(<ExtX>)</ExtX></ZoneX>

<ZoneX>: An integer in the range of {1, 2, 3, 4, 8, 10} declaring number of zones encountered across horizontal direction.

<TapX>: An integer in range {∅, 2, 3, 4, 8, 10} declaring the number of consecutive pixels in the horizontal direction that are outputted simultaneously from a zone. This field is omitted when all pixels are in the same column.

<ExtX>: A letter in the range of {∅, E, M} declaring the location of the pixels extractors in the horizontal direction. The value E indicates that pixel extractors are at both ends of the line. Value M indicates that pixel extractors are in the middle of the line. This field is omitted when all pixel extractors are all at the left end of each zone.

TapGeometryY is designated by <ZoneY>Y(<TapY>)(<ExtY>)</ExtY></ZoneY>

<ZoneY>: An integer in the range of {1, 2} declaring the number of zones encountered in the vertical direction.