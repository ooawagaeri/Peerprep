import CodeMirror from "@uiw/react-codemirror";
import { loadLanguage } from "@uiw/codemirror-extensions-langs";
import { ViewUpdate } from "@codemirror/view";
import * as Y from "yjs";
import React, { useEffect } from "react";
import { WebsocketProvider } from "y-websocket-peerprep";
import { yCollab } from "y-codemirror.next";
import { Language } from "../../types";

let providerSet = false;

type Props = {
  yText: Y.Text;
  provider: WebsocketProvider;
  undoManager: Y.UndoManager;
  nickname: string;
  selectedLang: Language;
  onCodeUpdate: (c: string, update: ViewUpdate) => void;
};

function Editor({
  yText,
  provider,
  undoManager,
  nickname,
  selectedLang,
  onCodeUpdate
}: Props) {
  useEffect(() => {
    if (!providerSet) {
      provider.awareness.setLocalStateField("user", {
        name: nickname,
        color: "#6eeb83",
        colorLight: "#6eeb8333"
      });
      providerSet = true;
    }

    return () => {};
  }, []);

  const lang: any = loadLanguage(selectedLang);

  return (
    <CodeMirror
      value=""
      height="100%"
      extensions={[lang, yCollab(yText, provider.awareness, { undoManager })]}
      onChange={onCodeUpdate}
    />
  );
}

export default Editor;
