|  **Category** | SequencerControl  |
| --- | --- |
|  **Level** | Recommended  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/Write  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | On Off  |

Controls if the sequencer configuration mode is active.

Possible values are:

- **Off:** Disables the sequencer configuration mode.
- **On:** Enables the sequencer configuration mode.

## 17.7 SequencerFeatureSelector

|  **Name** | SequencerFeatureSelector  |
| --- | --- |
|  **Category** | SequencerControl  |
|  **Level** | Recommended  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/Write  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | Device-Specific-Feature-List  |

Selects which sequencer features to control.

The feature lists all the features that can be part of a device sequencer set. All the device's sequencer sets have the same features.

Note that the name used in the enumeration must match exactly the device's feature name.

## 17.8 SequencerFeatureEnable

|  **Name** | SequencerFeatureEnable[SequencerFeatureSelector]  |
| --- | --- |
|  **Category** | SequencerControl  |
|  **Level** | Recommended  |