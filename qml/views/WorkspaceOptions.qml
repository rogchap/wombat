import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import QtQuick.Dialogs 1.1
import Qt.labs.platform 1.1

import "../."
import "../controls"

Modal {
    id: root

    readonly property var wc: mc.workspaceCtrl

    headerText: qsTr("Workspace")

    Column {
        spacing: 10

        TextField {
            id: txtAddr
            labelText: qsTr("gRPC server address:")
            placeholderText: "localhost:9090"
            text: wc.addr 
        }

        FileList {
            labelText: qsTr("Proto files:")
            actionButtonColor: Style.primaryColor
            actionButtonText: qsTr("Find *.proto files")

            model: wc.protoListModel

            onOpened: fdProtos.open()

            FolderDialog {
                id: fdProtos
                acceptLabel: qsTr("Find *.proto files")
                onAccepted: wc.findProtoFiles(folder)
            }
        }

        FileList {
            labelText: qsTr("Import proto paths:")

            model: wc.importListModel

            onOpened: fdImports.open()

            FolderDialog {
                id: fdImports
                onAccepted: wc.addImport(folder)
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
                let err = wc.processProtos()
                if (err) {
                    print(err)
                    return
                }
                err = wc.connect(txtAddr.text)
                if (err) {
                    print(err)
                    return
                }
                root.close()
            }
            
        }
        
    }

}
