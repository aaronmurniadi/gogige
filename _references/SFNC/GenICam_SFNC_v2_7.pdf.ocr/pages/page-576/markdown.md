|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

# 29 Tap Geometry Appendix

## 29.1 Motivations

This appendix defines standard names to uniquely identify the geometrical properties of most device's sensor taps.

For example, the initial release of the Camera Link® standard didn't include any information about the geometrical properties of taps.

Frame grabbers are able to reconstruct the image from multi-tap cameras on-the-fly.

Camera manufacturers should clearly specify what geometry(ies) is (are) required. Frame grabber manufacturers should also clearly specify what geometry(ies) is (are) supported.

The customer can then check the compatibility, and select the appropriate geometry for the camera and the frame grabber.

Considering the limited amount of cases, a unique name is assigned for each geometry.

## 29.2 Identifying the Geometrical Properties

### 29.2.1 Image Geometrical Properties

The relevant geometrical properties required for reconstructing the image:

- **Vantage point**: An enumerated value that specifies the position of the pixel with coordinate X=1, Y=1 in the scene.
{Top-Left, Top-Right, Bottom-Left, Bottom-Right}
Default is Top-Left.
- **ImageWidth**: An integer value declaring the image width expressed in pixels.
- **ImageHeight**: An integer value declaring the image height expressed in pixels. This parameter is irrelevant in case of line-scan or TDI cameras.
- **TapGeometry**: An enumerated type of parameter that summarizes the following properties for each tap:
  - XStart: X-coordinate of the first pixel column
  - YStart: Y-coordinate if the first pixel row
  - XEnd: X-coordinate of the last pixel column
  - YEnd: Y-coordinate of the last pixel row
  - XStep: Difference of X-coordinates between consecutive pixel columns; X-step is positive when X-coordinates are increasing along a row; it is negative otherwise.