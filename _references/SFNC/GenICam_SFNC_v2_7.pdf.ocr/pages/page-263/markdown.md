# 11 Encoder Control

This chapter lists all features that relates to control and monitoring of quadrature encoders. Quadrature encoders are also known as incremental, rotary and shaft encoders.

## 11.1 Quadrature Encoder usage model

This section describes the quadrature encoder usage model.

The Encoder Control interface connects to a rotary or linear quadrature encoder device which has output signals A and B used for tracking of positive and negative direction motion of the tracked axle or surface, respectively.

The Encoder Control interface uses a position counter driven by the A and B signals to track the position of the target in the presence of either a positive or a negative direction of motion.

The Encoder interface can generate a trigger signal to the sensor via TriggerSource while maintaining its counter value for tracking the position of the object. This means that reading the counter value gives the position and motion of the quadrature encoder (target object) from the time the position encoder started. Starting the position counter can be done with AcquisitionStart through the EncoderResetSource.

The number of input pulses from a quadrature encoder needed to generate Encoder Control output pulses, which can be used as a trigger source (i.e. the distance between samples along the encoder motion) is controlled via the EncoderDivider feature. The position counter runs continuously and its value represents the number of incoming pulses received since the last reset. Thus, this counter always reflects the true position of the object before the divider. The EncoderOutputMode feature is then used to control how the position counter generates its output signal.

Some quadrature encoders have an index output that indicates their per revolution position; however, this is not included explicitly in the Encoder Control interface. This signal can be handled through the standard I/O interface as a separate signal for example to reset of the Encoder Control position counter.

Quadrature encoder outputs often use differential signaling, but this conditioning is also handled by the generic I/O setup.

Note that the Encoder Control output can also be used to generate Events (See the EventSelector feature).