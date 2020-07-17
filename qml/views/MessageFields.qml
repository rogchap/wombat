import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.qmlmodels 1.0

import "../."
import "../controls"

ListView {
    id: root

    spacing: 10

    width: contentWidth
    height: contentHeight
    
    delegate: DelegateChooser {
        role: "delegate"

        DelegateChoice {
            roleValue: "text"
            DelegateTextField {
                onTextChanged: root.model.updateFieldValue(index, text) 
            }
        }
        
        DelegateChoice {
            roleValue: "bool"
            CheckBox {
                text: label
                checked: val == "true" ? true : false
                onCheckedChanged: root.model.updateFieldValue(index, checked) 

            }
        }

        DelegateChoice {
            roleValue: "message"
            DelegateMessageField {} 
        }

        DelegateChoice {
            roleValue: "textArea"
            TextAreaField {
                labelText: label
                hintText: type
                text: val
                onTextChanged: root.model.updateFieldValue(index, text) 
            }
        }

        DelegateChoice {
            roleValue: "enum"
            ComboBoxField {
                labelText: label
                model: enumListModel
                onDisplayTextChanged: root.model.updateFieldValue(index, displayText)
            }
        }

        DelegateChoice {
            roleValue: "text_repeated"
            DelegateRepeatedTextField {}
        }
    }
}
