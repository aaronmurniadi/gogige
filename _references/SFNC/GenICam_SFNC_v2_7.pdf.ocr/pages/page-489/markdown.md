|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- DisparityC_Linescan: Disparity 2.5D Depth map with varying B sampling. The distance is inversely proportional to the pixel (disparity) value. The B (Y) axis comes from the encoder chunk value.

24.60 ChunkScan3dCoordinateSystem

|  Name | ChunkScan3dCoordinateSystem  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | CartesianSphericalCylindricalDevice-specific  |

Returns the Coordinate System of the image included in the payload.

Possible values are:

- Cartesian: Default value. 3-axis orthogonal, right-hand X-Y-Z.
- Spherical: A Theta-Phi-Rho coordinate system.
- Cylindrical: A Theta-Y-Rho coordinate system.

24.61 ChunkScan3dCoordinateSystemReference

|  Name | ChunkScan3dCoordinateSystemReference  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Anchor Transformed  |