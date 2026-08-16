|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

### 6.3 Packed tag

Packed (“p”) is a common packing style where there is no bit spacing left between components (or possibly between successive pixels). The packed tag is followed by an optional number providing the number of bits the data is packed into (when it is not using all the available bits and padding bits are necessary) and by an optional bit order indicating if packing starts from lsb or msb. Empty bit must be padded to 0. The first component starts in byte 0.

#### 6.3.1 lsb Packed

lsb packed is the default packed mode and does not need to be explicitly specified after the 'p' indicator.

For lsb packed, the data is filled lsb first in the lowest address byte (byte 0) starting with the first component and continue in the lsb of byte 1 (and so on). Padding bits, if any, would thus be the msb's of the last byte after putting all the components. Padding bits are necessary when the "p" packing tag is followed by a number indicating to how many bits we need to align the pixel.

Note that in the following figures, we put byte 0 on the right to help illustrate the concept.

The following figure represents an example of a 3 color component pixel using 10 bits for each color component packed into a 32-bit data. The data is lsb packed; meaning byte 0 contains the least significant bits of the first color component. We start filling data with the lsb of byte 0 and continue with the lsb of byte 1 (and so on). The fact there is a 32 after the “p” packing tag indicates that padding bits are necessary to align the pixel to 32-bit in this example.

|   | byte 3 | byte 2 | byte 1 | byte 0  |
| --- | --- | --- | --- | --- |
|  ... | p p 9 . . . . 4N | 3 . . 0 9 . . 6N M | 5 . . . . 0 9 8M L | 7 . . . . . . 0L  |

Figure 6-7 :3 components in 10-bit lsb packed into 32-bit pixel (RGB10p32)

Notice that bits are put successively for each component with no spacing in-between, followed by the padding bits.

Here is another example typical for RGB565 lsb packed. Notice there is no number after the “p” packing tag hence no padding bit.

|   | byte 1 |   |   | byte 0  |   |   |
| --- | --- | --- | --- | --- | --- | --- |
|  ... | 4. | 0 | 5. | 3 | 2. | 0  |
|   | N | M |  | M | L |   |

Figure 6-8 :3 components lsb packed into 16-bit pixel (RGB565p)