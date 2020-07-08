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
            Label { text: "double: not implemented yet" }
        }
        
        DelegateChoice {
            roleValue: "TYPE_FLOAT"
            Label { text: "float: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_INT64"
            Label { text: "int64: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_UINT64"
            Label { text: "uint64: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_INT32"
            Label { text: "int32: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_FIXED64"
            Label { text: "fixed64: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_FIXED32"
            Label { text: "fixed32: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_BOOL"
            Label { text: "bool: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_STRING"
            TextField {
                labelText: label
                hintText: type.substring(5, type.length).toLowerCase()
                text: val
            }
        }

        DelegateChoice {
            roleValue: "TYPE_GROUP"
            Label { text: "group: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_MESSAGE"
            Item {
                height: msgPane.height + msgLabel.height + 10

                Label { id: msgLabel; text: label }

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
                }
            }
        }

        DelegateChoice {
            roleValue: "TYPE_BYTES"
            Label { text: "bytes: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_UINT32"
            Label { text: "uint32: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_ENUM"
            Label { text: "enum: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_SFIXED32"
            Label { text: "sfixed32: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_SFIXED64"
            Label { text: "sfixed64: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_SINT32"
            Label { text: "sint32: not implemented yet" }
        }

        DelegateChoice {
            roleValue: "TYPE_SINT64"
            Label { text: "sint64: not implemented yet" }
        }
    }

}
