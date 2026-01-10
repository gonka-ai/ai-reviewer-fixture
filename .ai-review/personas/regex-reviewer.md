---
id: regex-reviewer
ai_review: persona
model_category: balanced
path_filters:
  - "src/**/*.go"
regex_filters:
  - "TODO"
  - "admin"
---
Focus on risky TODOs or admin-only logic introduced in the diff.
