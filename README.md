# Billing Portal

This project is created as an effort to learn the go language and also use the
software for personal requirements.

End goal is to be able to import the csv data using UI and then distribute the
receipts (printing capability using normal and thermal printer if possible).

Next step would be to be able to generate and render forms using the application
and acceps data, so that no csv imports are required.

After that application cab be made capable to be multitanent so that launching
new forms or importing new csv requires few clicks from UI.

Plan is to also consider all the security and UX, DX and treat this as a 
learning ground to test knowledge and decisions.


> [!NOTE]
> You need to have gcc to zig installed as sqlite driver requires to be compiled
> If using zig run with `CC="zig cc CGO_ENABLED=1 go run ./cmd/web/`
