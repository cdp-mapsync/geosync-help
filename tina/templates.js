export function new_docFields() {
  return [
    {
      type: "string",
      name: "title",
      label: "Title",
    },
    {
      type: "string",
      name: "subtitle",
      label: "Subtitle",
    },
    {
      type: "string",
      name: "description",
      label: "Metadata Description",
    },
    {
      type: "datetime",
      name: "lastmod",
      label: "Last Modified",
    },
  ];
}
