|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- **Off**: The test mode is disabled. This feature has no effect and the device is streaming data normally according to its configuration. This option has to be the default after each boot of the device.
- **MultiPart**: The device streams data using multi-part payload format with at least one part in each payload. This option must be present if and only if the device supports the multi-part payload format.
- **GenDC**: The device streams data using GenDC payload format with at least one component in each payload. This option must be present if the device supports the GenDC payload format.

Note: If the underlying transport layer negotiation has failed to allow the device to enter multi-part or GenDC mode, it must not be possible to enable this test mode.