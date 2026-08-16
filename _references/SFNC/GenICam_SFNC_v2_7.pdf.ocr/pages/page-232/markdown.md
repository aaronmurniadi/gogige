|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

For a Digital I/O, when the full **I/O Control Block** is implemented, each physical **Line** (or pin) selected using **LineSelector** can be configured as Input or Output using **LineMode**. For an input or output Line, it is possible to read the Status of the Line with **LineStatus** and the incoming or outgoing signal can also be inverted using **LineInverter**. For an Output signal, the source of the signal is controlled using **LineSource** (See Figure 9-1).

For example:

/* Output an inverted pulse coming from the Timer 1 on the physical Line 2 of the device connector.

*/

LineSelector = Line2;
LineMode     = Output;
LineInverter = True;
LineSource   = Timer1Active;

/* Read the inverted Status of the physical Line 1. */

LineSelector  = Line1;
LineMode     = Input;
LineInverter = True;
CurrentStatus = LineStatus;

/* Output of the Exposure signal of each frame on the physical Line 2. */

LineSelector = Line2;
LineMode     = Output;
LineSource   = ExposureActive;

Note that all the features of an I/O control block are optional. Typically, an Input only line will report the **LineMode** as **Input** (read-only) and will implement only the **LineSelector**, **LineInverter** and **LineStatus** features (top half in Figure 9-1: I/O Control). An Output only line will report the **LineMode** as **Output** (read-only) and will implement only the **LineSelector**, **LineInverter** and **LineSource** features (bottom half of Figure 9-1: I/O Control). Even a hard-wired input or output line is just particular case where all the features are read-only.

The electrical format of the physical Line (TTL. LVDS, Opto-Coupled...) can be read or controlled (if supported) using **LineFormat**.

Note also that the Status of all the Lines can be monitored in one single access using **LineStatusAll**.

### UserOutput:

One possible source for Output lines is the User Output bit register.