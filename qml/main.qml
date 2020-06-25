import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.platform 1.1

ApplicationWindow {
    id: window
    visible: true
    title: "Courier" 
    minimumWidth: 1080
    minimumHeight: 720
    // flags: Qt.WindowStaysOnTopHint

    FolderDialog {
        id: fdProtoFolder
        onAccepted: {
            txtProtoFolder.text = folder
            mc.processProtos(folder)
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
                width: 200
            }
            ComboBox {
                id: cbMethodList
                textRole: "display"
                model: mc.methodList
                anchors.left: cbServiceList.right
                width: 200
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
