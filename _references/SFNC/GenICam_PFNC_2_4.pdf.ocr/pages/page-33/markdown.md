|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

#### Pixel construction rules for the “msb Unpacked” style

To construct the pixel stream:

1) Put the component in the most significant bits.
2) Pad the least significant bits to the nearest 8-bit boundary.
3) Start with the next component on the next 8-bit boundary.

### 6.2 Cluster marker

The cluster marker (“c”) is only allowed for single-component/monochrome pixel formats. It is used to regroup a given number of single-component/monochrome pixels into one multi-component pixel. This facilitates the re-use of some multi-component/color pixel packing style concepts for single-component/monochrome pixels.

The cluster marker is immediately followed by a number indicating the number of single-component/monochrome pixels that are grouped into the cluster.

Ex: c2 = 2 single-component/monochrome pixels in the cluster

c3 = 3 single-component/monochrome pixels in the cluster (which makes the cluster similar to RGB format)

A cluster marker is only required to remove a possible ambiguity with the pixel format name, typically when the number of bits (including padding) is not a multiple of the number of single-component/monochrome pixels in the cluster. In general, the cluster marker should be avoided as it clouds the pixel name and makes it less friendly.

When the cluster marker is used, then the packed or grouped style must consider the cluster of single-component/monochrome pixels as one multi-component pixel. This directly impacts the number immediately following those 2 tags which must now represent the number of bits for the cluster.

The following figure illustrates a scenario where 3 monochrome pixels are regrouped into one 3-component pixel. This 3-component pixel is then lsb packed to 32 bits, leaving 2 padding bits in the msb's position.

|   | byte 3 | byte 2 | byte 1 | byte 0  |
| --- | --- | --- | --- | --- |
|  ... | p p 9 . . . . 4L3 | 3 . . 0 9 . . 6L3 L2 | 5 . . . . 0 9 8L2 L1 | 7 . . . . . . 0L1  |

Figure 6-6 : 10-bit monochrome pixel lsb packed into 32 bits (Mono10c3p32)