|  ComponentEnable[RegionSelector][ComponentSelector] | O | IBoolean | R/(W) | - | B | Controls if the selected component streaming is active.  |
| --- | --- | --- | --- | --- | --- | --- |
|  ComponentIDValue[ComponentSelector] | O | IInteger | R | - | E | Returns a unique Identifier value that corresponds to type of the component selected by ComponentSelector.  |
|  GroupSelector | O | IEnumeration | R/W | - | B | Selects a Group of component to control or inquire.  |
|  GroupIDValue[GroupSelector] | O | IInteger | R | - | E | Returns a unique Identifier value corresponding to the selected Group of Components.  |
|  ImageComponentSelector | O | IEnumeration | R/W | - | I | This feature is deprecated (See ComponentSelector).  |
|  ImageComponentEnable[RegionSelector][ComponentSelector] | O | IBoolean | R/(W) | - | I | This feature is deprecated (See ComponentEnable).  |
|  Width[RegionSelector] | R | IInteger | R/(W) | - | B | Width of the image provided by the device (in pixels).  |
|  Height[RegionSelector] | R | IInteger | R/(W) | - | B | Height of the image provided by the device (in pixels).  |
|  OffsetX[RegionSelector] | R | IInteger | R/W | - | B | Horizontal offset from the origin to the region of interest (in pixels).  |
|  OffsetY[RegionSelector] | R | IInteger | R/W | - | B | Vertical offset from the origin to the region of interest (in pixels).  |
|  LinePitchEnable[RegionSelector] | R | IBoolean | R/W | - | E | This feature controls whether the LinePitch feature is writable.  |
|  LinePitch[RegionSelector] | R | IInteger | R/W | B | E | Total number of bytes between the starts of 2 consecutive lines.  |
|  BinningSelector | O | IEnumeration | R/(W) | - | E | Selects which binning engine is controlled by the BinningHorizontal and BinningVertical features.  |
|  BinningHorizontalMode[BinningSelector] | O | IEnumeration | R/(W) | - | E | Sets the mode to use to combine horizontal photo-sensitive cells together when BinningHorizontal is used.  |
|  BinningHorizontal[BinningSelector] | O | IInteger | R/W | - | E | Number of horizontal photo-sensitive cells to combine together.  |
|  BinningVerticalMode[BinningSelector] | O | IEnumeration | R/(W) | - | E | Sets the mode to use to combine vertical photo-sensitive cells together when BinningVertical is used.  |
|  BinningVertical[BinningSelector] | O | IInteger | R/W | - | E | Number of vertical photo-sensitive cells to combine together.  |
|  DecimationHorizontalMode | O | IEnumeration | R/(W) | - | E | Sets the mode used to reduce the horizontal resolution when DecimationHorizontal is used.  |
|  DecimationHorizontal | O | IInteger | R/W | - | E | Horizontal sub-sampling of the image.  |