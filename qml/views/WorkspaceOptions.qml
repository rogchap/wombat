import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import Qt.labs.platform 1.1

import "../."
import "../controls"

Modal {
    id: root

    headerText: qsTr("Workspace")

    Column {
        spacing: 10

        TextField {
            id: txtAddr
            labelText: qsTr("gRPC server address:")
            placeholderText: "localhost:9090"
            text: mc.addr 
        }

        FileList {
            labelText: qsTr("Proto files:")
            actionButtonColor: Style.primaryColor
            actionButtonText: qsTr("Find *.proto files")

            model: mc.protoFilesList

            onOpened: fdProtos.open()

            FolderDialog {
                id: fdProtos
                acceptLabel: qsTr("Find *.proto files")
                onAccepted: mc.findProtoFiles(folder)
            }
        }

        FileList {
            labelText: qsTr("Import proto paths:")

            model: mc.protoImportsList

            onOpened: fdImports.open()

            FolderDialog {
                id: fdImports
                onAccepted: mc.addImport(folder)
            }
        }

        Rectangle {
            height: 1
            width: parent.width
            color: Style.borderColor
        }

        Button {
            anchors.right: parent.right
            bgColor: Style.accentColor3

            text: qsTr("Connect")

            onClicked: {
                // TODO: handle any errors
                mc.processProtos("","")
                mc.updateAddr(txtAddr.text)
                root.close()
            }
        }
        
    }

}
