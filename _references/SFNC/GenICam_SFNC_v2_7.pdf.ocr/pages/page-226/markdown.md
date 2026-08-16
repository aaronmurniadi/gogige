## 8 Color Transformation Control

The Color Transformation chapter describes all features related to color Transformations in the device.

The Color Transformation is a linear operation taking as input a triplet of Components (C0, C1, C2) for a color pixel (Typically: R_in, G_in, B_in representing a RGB color pixel). This triplet is first multiplied by a 3x3 matrix and then added to an offset triplet.

The equation is given by:

$$\left( \begin{array}{c} \text{RC0}_{out} \\ \text{G} \\ \text{C1}_{out} \\ \text{B} \\ \text{C2}_{out} \end{array} \right) = \left( \begin{array}{ccc} Gain00 & Gain01 & Gain02 \\ Gain10 & Gain11 & Gain12 \\ Gain20 & Gain21 & Gain22 \end{array} \right) \left( \begin{array}{c} C0_{in} \\ C1_{in} \\ C2_{in} \end{array} \right) + \left( \begin{array}{c} Offset0 \\ Offset1 \\ Offset2 \end{array} \right)$$

$$\text{Equivalent: } \left( \begin{array}{c} R_{out} \\ G_{out} \\ B_{out} \end{array} \right) = \left( \begin{array}{ccc} RR & RG & RB \\ GR & GG & GB \\ BR & BG & BB \end{array} \right) \left( \begin{array}{c} R_{in} \\ G_{in} \\ B_{in} \end{array} \right) + \left( \begin{array}{c} R_{offset} \\ G_{offset} \\ B_{offset} \end{array} \right)$$

The descriptions below assume RGB to RGB transformation:

|  Where | C0_{in} is the first component of the incoming pixel  |
| --- | --- |
|   | C1_{in} is the second component of the incoming pixel  |
|   | C2_{in} is the third component of the incoming pixel  |
|   | Gain00 is the red contribution to the red pixel (multiplicative factor)  |
|   | Gain01 is the green contribution to the red pixel (multiplicative factor)  |
|   | Gain02 is the blue contribution to the red pixel (multiplicative factor)  |
|   | Gain10 is the red contribution to the green pixel (multiplicative factor)  |
|   | Gain11 is the green contribution to the green pixel (multiplicative factor)  |
|   | Gain12 is the blue contribution to the blue pixel (multiplicative factor)  |
|   | Gain20 is the red contribution to the blue pixel (multiplicative factor)  |
|   | Gain21 is the green contribution to the blue pixel (multiplicative factor)  |
|   | Gain22 is the blue contribution to the blue pixel (multiplicative factor)  |
|   | Offset0 is the red offset  |
|   | Offset1 is the green offset  |