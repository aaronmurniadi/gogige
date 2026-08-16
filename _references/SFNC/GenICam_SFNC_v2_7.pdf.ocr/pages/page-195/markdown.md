|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- **ExposureMode** can be **Timed** to have a timed exposure and allow programming the duration using the **ExposureTime** or **ExposureAuto** features.

For example to have a fixed exposure time of 1 millisecond, use the following pseudo code:

Camera.ExposureMode = Timed;
Camera.ExposureTime = 1000;

- **ExposureMode** can be **TriggerWidth** to use the width of the current Frame or Line trigger signal(s) to control exposure duration.
- **ExposureMode** can be **TriggerControlled** to use one or more trigger signal(s) to control the exposure duration independently from the current Frame or Line triggers (See **ExposureStart**, **ExposureEnd** and **ExposureActive** of the **TriggerSelector** feature).

For example: To use 2 hardware triggers respectively starting and stopping the Exposure, use the following pseudo code:

Camera.ExposureMode = TriggerControlled;
Camera.TriggerSelector = ExposureStart;
Camera.TriggerMode = On;
Camera.TriggerSource = Line1;
Camera.TriggerSelector = ExposureEnd;
Camera.TriggerMode = On;
Camera.TriggerSource = Line2;