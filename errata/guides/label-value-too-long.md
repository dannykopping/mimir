This non-critical error occurs when Mimir receives a write request that contains a series with a label value whose length exceeds the configured limit.
The limit protects the system’s stability from potential abuse or mistakes. To configure the limit on a per-tenant basis, use the `-validation.max-length-label-value` option.

> **Note**: Invalid series are skipped during the ingestion, and valid series within the same request are ingested.