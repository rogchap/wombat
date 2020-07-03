import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import Qt.labs.platform 1.1


ApplicationWindow {
    id: window
    visible: true
    title: "Courier" 
    minimumWidth: 1200
    minimumHeight: 820
    // flags: Qt.WindowStaysOnTopHint


    FolderDialog {
        id: fdFindProtos
        acceptLabel: "Find *.proto files"
        onAccepted: mc.findProtoFiles(folder)
    }

    FolderDialog {
        id: fdImportFolder
        acceptLabel: "Select"
        onAccepted: mc.addImport(folder)
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
                anchors.fill: parent
                spacing: 5

                TextField {
                    id: txtServer
                    width: parent.width
                    text: "localhost:10000"
                    placeholderText: "grpc server URL" 
                }
                Label {
                    text: "Proto files:"
                    leftPadding: 5
                }
                ScrollView {
                    ScrollBar.horizontal.policy: ScrollBar.AlwaysOff
                    width: parent.width
                    height: 150
                    clip: true
                    background: Rectangle {
                        color: "white"
                        border {
                            color: "lightgray"
                            width: 1
                        }
                    }

                    ListView {
                        anchors.fill: parent
                        model: mc.protoFilesList
                        delegate: Label {
                            width: parent.width
                            text: display
                            elide: Text.ElideLeft
                        }
                    }
                }
                Row {
                    width: parent.width
                    layoutDirection: Qt.RightToLeft
                    rightPadding: 5

                    Button {
                        text: "Find *.proto files"
                        onClicked: fdFindProtos.open()
                    }
                }

                Label {
                    text: "Import paths:"
                    leftPadding: 5
                }

                ScrollView {
                    ScrollBar.horizontal.policy: ScrollBar.AlwaysOff
                    width: parent.width
                    height: 150
                    clip: true
                    background: Rectangle {
                        color: "white"
                        border {
                            color: "lightgray"
                            width: 1
                        }
                    }
                    ListView {
                        anchors.fill: parent
                        model: mc.protoImportsList
                        delegate: Label {
                            width: parent.width
                            text: display
                            elide: Text.ElideLeft
                        }
                    }
                }
                Row {
                    width: parent.width
                    layoutDirection: Qt.RightToLeft
                    rightPadding: 5

                    Button {
                        text: "Select folder"
                        onClicked: fdImportFolder.open()
                    }
                }

                Item {
                    width: parent.width
                    height: 10
                    // spacer    
                }
                Button {
                    text: "Connect"
                    width: parent.width
                    onClicked: mc.processProtos("","")
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
            Rectangle {
                id: inputContainer
                anchors.fill: parent
                anchors.topMargin: cbServiceList.height
                color: "lightgray"
                ColumnLayout {
                    anchors.fill: parent
                    Label {
                        id: lblInput
                        text: mc.input.label
                    }
                    ListView {
                        Layout.fillHeight: true
                        spacing: 5
                        model: mc.input
                        delegate: Row {
                            Label {
                                text: label
                            }
                            TextField {
                                text: val
                                selectByMouse: true
                                selectionColor: "lightgray"
                                onTextChanged: mc.input.updateFieldValue(index, text) 
                            }
                        }
                    }
                }

            }
        }
        Rectangle {
            implicitWidth: 300
            color: "lightgreen"
            TextArea {
                text: mc.output
                anchors.fill: parent
                wrapMode: TextEdit.WordWrap
                selectByMouse: true
                selectionColor: "lightgray"
            }
        }
    }
}
