<script>
  import init from "/src/example.wasm?init";
  import "/src/wasm_exec.js";

  // Init WASM
  const go = new globalThis.Go();
  init(go.importObject).then((instance) => go.run(instance));

  let template = "Hi {{ .where }}";
  let jsonData = `{"where": "there 🐣"}`;
  let templateErr = null;
  let jsonErr = null;
  let result = "Hi there 🐣";

  const handleTemplateChange = async (e) => {
    template = e.target.value.trim();
    await updateResult("template");
  };

  const handleJsonDataChange = async (e) => {
    jsonData = e.target.value.trim();
    await updateResult("json");
  };

  const updateResult = async (updatedValue) => {
    const response = await globalThis.parse(template, jsonData);
    const { templateError, jsonError, parsed } = response;

    console.log(templateError, jsonError, parsed);
    templateErr = templateError === "" ? null : templateError;
    jsonErr = jsonError === "" ? null : jsonError;
    result = parsed;
  };
</script>

<main>
  <section class="inputs">
    <textarea
      id="template"
      placeholder="template"
      class={templateErr && "error"}
      on:keyup={handleTemplateChange}>{template}</textarea
    >
    <textarea
      id="json"
      class={jsonErr && "error"}
      placeholder="json"
      on:keyup={handleJsonDataChange}>{jsonData}</textarea
    >
  </section>
  <p class={(jsonErr || templateErr) && "error"}>
    {jsonErr || templateErr || result}
  </p>
</main>
