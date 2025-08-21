import "@typespec/http";

using Http;
@service(#{ title: "Widget Service" })
namespace DemoService;

@doc("Widget is a nice thing to have")
model Widget {
  id: string;
  weight: int32;
  color: "red" | "blue";
}

model WidgetList {
  items: Widget[];
}

@error
model Error {
  code: int32;
  message: string;
}

model AnalyzeResult {
  id: string;
  analysis: string;
}

model WidgetCreateRequest is OmitProperties<Widget, "id">;
model WidgetUpdateRequest is MergePatchUpdate<OmitProperties<Widget, "id">>;

@route("/widgets")
@tag("Widgets")
interface Widgets {
  /** List widgets */
  @get list(): WidgetList | Error;
  /** Read widgets */
  @get read(@path id: string): Widget | Error;
  /** Create a widget */
  @post create(@body body: WidgetCreateRequest): Widget | Error;
  /** Update a widget */
  @patch update(
    @path id: string,
    @body body: WidgetUpdateRequest,
  ): Widget | Error;
  /** Delete a widget */
  @delete delete(@path id: string): void | Error;

  /** Analyze a widget */
  @route("{id}/analyze") @post analyze(@path id: string): AnalyzeResult | Error;
}
