|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 1.1 Conventions

### Feature Name and Interface

According to the GenICam standard, all the public features of a device must be included in the GenICam XML file and must use the SFNC Name and Interface type for those features if they exist. Other vendor specific or specialized features not mapping to existing SNFC features can be included but must be located in a vendor specific namespace in the GenICam XML and may use a vendor specific name.

This document lists for each feature, the Name and Interface type that must be used.

### Feature Category

With the GenICam standard, each feature should be included in a "Category". The Category element defines in which group of features, the feature will be located.

The Category does not affect the functionality of the features but is used by the GUIs to group the features when displaying them. The purpose is mainly to insure that the GUI can present features in a more organized way.

This document lists for each feature, a recommended Category that should be used.

### Feature Level

In this document, features are tagged according to the following requirement levels:

- M: Mandatory - Must be implemented to achieve compliance with the GenICam standard.
- R: Recommended - This feature adds important aspects to the use case and must respect the naming convention if used.
- O: Optional - This feature is less critical. Nevertheless, it is considered and must respect the naming convention if used.

For additional details about the mandatory features specific to a particular transport layers, please refer to the text of those standards.

### Feature Visibility

According to the GenICam standard each feature can be assigned a "Visibility". The Visibility defines the type of user that should get access to the feature. Possible values are: Beginner, Expert, Guru and Invisible. The latter is required to make features accessible from the API, but invisible in the GUI.

The visibility does not affect the functionality of the features but is used by the GUI to decide which features to display based on the current user level. The purpose is mainly to insure that the GUI is not cluttered with information that is not intended at the current user level.

The following criteria have been used for the assignment of the recommended visibility: