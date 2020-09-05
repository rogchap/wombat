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
                rhs: true
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
                currentIndex: enumListModel.idxForVal(val)
                onDisplayTextChanged: root.model.updateFieldValue(index, enumListModel.valAt(currentIndex))
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
                        text: value
                        onTextChanged: valueListModel.editValueAt(index, text)
                    }

                    CrossButton {
                        color: Style.bgColor3
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
                        rhs: true
                        text: label
                        anchors.left: parent.left
                        anchors.leftMargin: 16
                        checked: value == "true" ? true : false

                        onCheckedChanged: valueListModel.editValueAt(index, checked) 
                    }

                    CrossButton {
                        color: Style.bgColor3
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
                        currentIndex: enumListModel.idxForVal(value)
                        onDisplayTextChanged: valueListModel.editValueAt(index, enumListModel.valAt(currentIndex))
                    }

                    CrossButton {
                        color: Style.bgColor3
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
                        text: value
                        onTextChanged: valueListModel.editValueAt(index, text)
                    }

                    CrossButton {
                        color: Style.bgColor3
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
                        onClicked: valueListModel.remove(index)
                    }

                }
            }
        }
    }
}
