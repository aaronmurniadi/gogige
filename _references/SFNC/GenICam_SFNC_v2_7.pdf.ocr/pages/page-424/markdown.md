|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

The Optic features intend to support different use cases. Some values might seem to be duplicates (e.g. Aperture and ApertureStepper) but in reality, are different.

The optical feature might be implemented using a stepper motor or a digital-to-analog converter, which has a defined start, stop and number of steps. These steps might have a linear, logarithmic or even more complex impact on the optical value. It is common that several controls are implemented in one device and that changing one influences the optical interpretation of the other. All stepper values are defined as being in the native resolution of the device.

For example, if a device has two steppers, one for Focal length and one for Aperture. If the Focal length is modified, the stepper value of the aperture will stay the same while the optical interpretation of the Aperture might change.

The Stepper value registers will control these steps, and every change in a register will lead to a change in the mechanical position or digital-to-analog converter input. Rational for including real-world (human-interpretable) values:

1. A user setting up the device profits from directly interpretable values, e.g. object-sensor-distance in millimeters, the aperture in f/#. Using these settings, the result will stay the same when changing lenses.
2. The optical interpretation enables calculation of optical features, e.g. to keep the brightness of an image identical when changing the exposure time, the f/# value of aperture can be used.

Rational for including the Stepper values:

1. A closed control loop on a real world value might change another real-world-value as well.
2. A change of Focal length might change the interpretation of the Aperture, but not the stepper value. Two control loops on the real-world value would constantly work against each other, while two control loops on the stepper value will not change the other stepper value