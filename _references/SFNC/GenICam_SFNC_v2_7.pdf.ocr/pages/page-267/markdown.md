|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

the motion will increment or decrement the position counter. The user can select Up or Down counting for the Position and Direction modes.

### Encoder Status:

The Encoder Status reflects if the encoder is active or not, and can be used for instance to create an event if the encoder stops receiving input pulses for a predefined timeout time, or as "wakeup" when it restarts its motion after a timeout.

### Encoder real world translation:

The encoder information can convey real world distance and displacement if the step size of the encoder is used. For 3D cameras the Scan3dCoordinateScale can be used for this purpose.

### Encoder counter wrap around:

If the position counter is running continuously it will wrap around between its maximum and minimum values if moving in positive and negative directions.

The position for wrap-around should be readable for the application through the Min and Max values of the EncoderValue feature.

This means that the application using the EncoderValue of a scan to track large displacements needs to use the wrap around limits and use corrective calculations when calculating the true displacement. For instance, when a sudden jump from ~encoder max to ~encoder min is observed in the encoder value the receiver can easily correct the true displacement by adding a full position counter (encodermax-encodermin+1) to the result. As long as the range of the position counter is much larger than the delta between individual encoder values for scans this method works well.

### Example of Encoder setup:

// Encoder Setup
EncoderSelector = Encoder0;
EncoderSourceA[Encoder0] = Line0;
EncoderSourceB[Encoder0] = Line1;
EncoderMode[Encoder0] = FourPhase;
EncoderDivider[Encoder0] = 10;
EncoderOutputMode[Encoder0] = PositionUp;

To set the destination output Line of the Encoder module, see Encoder entries of the LineSource feature.