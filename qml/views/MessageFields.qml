import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.qmlmodels 1.0

import "../."
import "../controls"

ListView {
    id: root

    spacing: 10

    width: parent.width
    height: contentHeight
    
    delegate: DelegateChooser {
        role: "delegate"

        DelegateChoice {
            roleValue: "text"
            DelegateTextField {}
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
            Item {
                height: msgPane.height + msgLabel.height + 10

                Label {
                    id: msgLabel
                    text: label
                    anchors.left: parent.left
                    anchors.leftMargin: 5
                }

                Label { 
                    anchors.left: msgLabel.right
                    anchors.leftMargin: 10
                    color: Qt.darker(Style.textColor3, 1.6)
                    text: message.label
                }

                Pane {
                    id: msgPane

                    width: msgLoader.width
                    height: msgLoader.height
                    anchors.top: msgLabel.bottom

                    Loader {
                        id: msgLoader

                        source: "MessageFields.qml"
                        onLoaded: {
                            item.model = message
                        }
                    }

                    Rectangle {
                        width: 1
                        height: msgPane.height + 5
                        color: Style.accentColor
                        anchors.left: parent.left
                        anchors.top: parent.top
                        anchors.leftMargin: -7
                        anchors.topMargin: -5
                    }
                }
            }
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
    }
}
