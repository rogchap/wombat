import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.platform 1.1

import "../."
import "../controls"

Pane {
    id: root

    property alias addr: txtAddr.text

    padding: 0
    topPadding: 10

    Column {
        spacing: 10

        TextField {
            id: txtAddr
            labelText: qsTr("gRPC server address:")
            placeholderText: "localhost:9090"
            text: wc.options.addr 
        }

        CheckBox {
            text: qsTr("Use reflection to determine the RPC schema")
        }

        Row {
            spacing: 10

            FileList {
                labelText: qsTr("Proto source file(s):")
                actionButtonColor: Style.primaryColor
                actionButtonText: qsTr("Find *.proto files")

                model: wc.options.protoListModel

                onOpened: fdProtos.open()
                onCleared: wc.options.clearProtoList()

                FolderDialog {
                    id: fdProtos
                    acceptLabel: qsTr("Find *.proto files")
                    onAccepted: wc.findProtoFiles(folder)
                }
            }

            FileList {
                labelText: qsTr("Import proto path(s):")

                model: wc.options.importListModel

                onOpened: fdImports.open()
                onCleared: wc.options.clearImportList()

                FolderDialog {
                    id: fdImports
                    onAccepted: wc.addImport(folder)
                }
            }

        }
    }
}
