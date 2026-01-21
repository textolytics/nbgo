
#!/usr/bin/env bash
set -e

cat <<'PYCODE'
# ======================================================
# SCHEMA DISCOVERY & VALIDATION (JUPYTER / IPYNB SAFE)
# ======================================================

print("=" * 70)
print("SCHEMA DISCOVERY: Discovering provider schemas")
print("=" * 70)

providers = [
    "CProvider",
    "RustProvider",
    "GoProvider",
    "PythonProvider",
    "GateGateway",
    "FreedxGateway",
]

async def main():
    print(f"\nDiscovering {len(providers)} providers...\n")

    # ✅ Correct async usage for ipynb (NO asyncio.run)
    result = await discovery.discover_providers(providers)

    print(f"✓ Providers found: {result['providers_found']}")
    print(f"✓ Schemas generated: {result['schemas_generated']}")
    print(f"✓ Scan completed at: {result['scan_time']}")

    print("\n\nDISCOVERED SCHEMAS:")
    print("-" * 70)

    for i, schema in enumerate(discovery.list_schemas(), 1):
        print(f"  {i}. {schema.name:20} v{schema.version}")
        print(f"     Fields: {', '.join(f.name for f in schema.fields)}")

    print("\n\nSCHEMA VALIDATION TEST:")
    print("-" * 70)

    test_data = {
        "name": "MyProvider",
        "version": "1.0.0",
        "enabled": True,
    }

    valid, errors = discovery.validate_schema("schema_CProvider", test_data)

    print(f"Data: {test_data}")
    print(f"Validation: {'✓ VALID' if valid else '✗ INVALID'}")

    if errors:
        for error in errors:
            print(f"  Error: {error}")

# ✅ Single correct execution point for notebooks
await main()
PYCODE
