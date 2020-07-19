import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.qmlmodels 1.0

import "../."
import "../controls"

ListView {
    id: root

    spacing: 10

    implicitWidth: contentWidth
    implicitHeight: contentHeight
    
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
            DelegateRepeatedField {
                delegate: Item {
                    height: textField.height
                    width: parent.width

                    DelegateTextField {
                        id: textField
                        hintText: type 
                        labelLeftMargin: 21
                    }

                    CrossButton {
                        color: Style.bgColor3
                        rotation: 45
                        onClicked: valueListModel.remove(index)
                    }
                }
            }
        }

        DelegateChoice {
            roleValue: "bool_repeated"
            DelegateRepeatedField {
                delegate: Item{
                    height: checkboxField.height
                    width: parent.width

                    CheckBox {
                        id: checkboxField
                        text: label
                        anchors.left: parent.left
                        anchors.leftMargin: 16
                    }

                    CrossButton {
                        color: Style.bgColor3
                        rotation: 45
                        onClicked: valueListModel.remove(index)
                        anchors.verticalCenter: parent.verticalCenter
                    }

                }
            }
        }

        DelegateChoice {
            roleValue: "enum_repeated"
            DelegateRepeatedField {
                delegate: Item{
                    height: comboField.height
                    width: parent.width

                    ComboBoxField {
                        id: comboField
                        labelText: label
                        labelLeftMargin: 21
                        model: enumListModel
                    }

                    CrossButton {
                        color: Style.bgColor3
                        rotation: 45
                        onClicked: valueListModel.remove(index)
                    }

                }
            }
        }

        DelegateChoice {
            roleValue: "textArea_repeated"
            DelegateRepeatedField {
                delegate: Item{
                    height: areaField.height
                    width: parent.width

                    TextAreaField {
                        id: areaField
                        labelText: label
                        labelLeftMargin: 21
                        hintText: type
                    }

                    CrossButton {
                        color: Style.bgColor3
                        rotation: 45
                        onClicked: valueListModel.remove(index)
                    }

                }
            }
        }

        DelegateChoice {
            roleValue: "message_repeated"
            DelegateRepeatedField {
                hintText: "repeated "+ message.label
                delegate: Item{
                    height: msgField.height
                    width: parent.width

                    DelegateMessageField {
                        id: msgField
                        msgOverride: msgValue
                        labelLeftMargin: 21
                        labelColor: Style.accentColor2
                    }

                    CrossButton {
                        color: Style.bgColor3
                        rotation: 45
                        onClicked: valueListModel.remove(index)
                    }

                }
            }
        }
    }
}
