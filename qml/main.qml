import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.platform 1.1

ApplicationWindow {
    id: window
    visible: true
    title: "Courier"
    minimumWidth: 800
    minimumHeight: 450

    Shortcut {
        sequence: "r"
        onActivated: {
            window.close();
            eCtx.reload();
        }
    }

    FolderDialog {
        id: fdProtoDir
        onAccepted: txtProtoFolder.text = folder
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
                    width: parent.width
                    anchors.top: txtServer.bottom
                    anchors.left: parent.left
                    TextField {
                        id: txtProtoFolder
                        placeholderText: "folder to proto files"
                        anchors.right: btnProtoOpen.left
                        anchors.left: parent.left
                    }
                    Button {
                        id: btnProtoOpen
                        text: "open"
                        anchors.right: parent.right
                        onClicked: fdProtoDir.open()
                    }
                }
            }
        }
        Rectangle {
            id: centerItem
            SplitView.minimumWidth: 50
            SplitView.fillWidth: true
            color: "lightgray"
            Label {
                text: "View 2"
                anchors.centerIn: parent
            }
        }
        Rectangle {
            implicitWidth: 200
            color: "lightgreen"
            Label {
                text: "View 3"
                anchors.centerIn: parent
            }
        }
    }
}
