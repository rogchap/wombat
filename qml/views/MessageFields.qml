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
        role: "type"

        DelegateChoice {
            roleValue: "TYPE_DOUBLE"
            DelegateTextField {
                validator: DoubleValidator {}
            }
        }
        
        DelegateChoice {
            roleValue: "TYPE_FLOAT"
            DelegateTextField {
                validator: DoubleValidator {}
            }
        }

        DelegateChoice {
            roleValue: "TYPE_INT64"
            DelegateTextField {
                validator: RegExpValidator { regExp: /^-?\d+/ }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_UINT64"
            DelegateTextField {
                validator: RegExpValidator { regExp: /\d+/ }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_INT32"
            DelegateTextField {
                validator: IntValidator { bottom: -2147483648; top: 2147483647 }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_FIXED64"
            DelegateTextField {
                validator: RegExpValidator { regExp: /^-?\d+/ }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_FIXED32"
            DelegateTextField {
                validator: IntValidator { bottom: -2147483648; top: 2147483647 }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_BOOL"
            CheckBox {
                text: label
                checked: val == "true" ? true : false
                onCheckedChanged: root.model.updateFieldValue(index, checked) 

            }
        }

        DelegateChoice {
            roleValue: "TYPE_STRING"
            DelegateTextField {}
        }

        DelegateChoice {
            roleValue: "TYPE_GROUP"
            Label { text: "group: not supported" }
        }

        DelegateChoice {
            roleValue: "TYPE_MESSAGE"
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
            roleValue: "TYPE_BYTES"
            TextAreaField {
                labelText: label
                hintText: type.substring(5, type.length).toLowerCase()
                text: val
                onTextChanged: root.model.updateFieldValue(index, text) 
            }
        }

        DelegateChoice {
            roleValue: "TYPE_UINT32"
            DelegateTextField {
                validator: RegExpValidator { regExp: /^-?\d+/ }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_ENUM"
            Label { text: "enum: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_SFIXED32"
            DelegateTextField {
                validator: IntValidator { bottom: -2147483648; top: 2147483647 }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_SFIXED64"
            DelegateTextField {
                validator: RegExpValidator { regExp: /^-?\d+/ }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_SINT32"
            DelegateTextField {
                validator: IntValidator { bottom: -2147483648; top: 2147483647 }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_SINT64"
            DelegateTextField {
                validator: RegExpValidator { regExp: /^-?\d+/ }
            }
        }
    }

}
