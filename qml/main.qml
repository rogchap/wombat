import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.platform 1.1

ApplicationWindow {
    id: window
    visible: true
    title: "Courier" 
    minimumWidth: 1200
    minimumHeight: 820
    // flags: Qt.WindowStaysOnTopHint

    FolderDialog {
        id: fdProtoFolder
        onAccepted: {
            txtProtoFolder.text = folder
            mc.processProtos(txtImportFolder.text, folder)
        }
    }

    FolderDialog {
        id: fdImportFolders
        onAccepted: {
            txtImportFolder.text = folder
        }
    }

    SplitView {
        anchors.fill: parent
        orientation: Qt.Horizontal
        handle: Rectangle {
            id: handle
            implicitWidth: 1
        }

        Rectangle {
            implicitWidth: 300
            SplitView.maximumWidth: 400
            color: "lightblue"
            Column {
                width: parent.width
                TextField {
                    id: txtServer
                    width: parent.width
                    placeholderText: "grpc server URL" 
                }
                Item {
                    id: protoFolderContainer
                    width: parent.width
                    anchors.top: txtServer.bottom
                    anchors.left: parent.left
                    TextField {
                        id: txtProtoFolder
                        readOnly: true
                        placeholderText: "folder to proto files"
                        anchors.right: btnProtoOpen.left
                        anchors.left: parent.left
                    }
                    Button {
                        id: btnProtoOpen
                        text: "open"
                        anchors.right: parent.right
                        onClicked: fdProtoFolder.open()
                    }
                }
                Item {
                    width: parent.width
                    anchors.topMargin: 40
                    anchors.top: protoFolderContainer.bottom
                    anchors.left: parent.left
                    TextField {
                        id: txtImportFolder
                        readOnly: true
                        placeholderText: "folder to proto imports"
                        anchors.right: btnImportOpen.left
                        anchors.left: parent.left
                    }
                    Button {
                        id: btnImportOpen
                        text: "open"
                        anchors.right: parent.right
                        onClicked: fdImportFolders.open()
                    }
                }
            }
        }
        Rectangle {
            id: centerItem
            SplitView.minimumWidth: 50
            SplitView.fillWidth: true
            color: "lightgray"
            ComboBox {
                id: cbServiceList
                textRole: "display"
                model: mc.serviceList
                width: 300
                onActivated: {
                    console.log("here")
                    mc.serviceChanged(displayText)
                }
            }
            ComboBox {
                id: cbMethodList
                textRole: "display"
                model: mc.methodList
                anchors.left: cbServiceList.right
                width: 200
            }
            Button {
                id: btnSend
                text: "send"
                anchors.left: cbMethodList.right
                onClicked: mc.send(txtServer.text, cbServiceList.displayText, cbMethodList.displayText)
            }
        }
        Rectangle {
            implicitWidth: 300
            color: "lightgreen"
            Label {
                text: mc.output
                anchors.centerIn: parent
            }
        }
    }
}
