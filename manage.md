# 📌 NOTION / TRELLO ROADMAP STRUCTURE

Format board management.

---

## 📂 BOARD: PARKING SYSTEM DEV

---

## 🟦 COLUMN: BACKLOG

- Design ERD
- Design flow
- Role mapping
- Pricing rules
- Security rules

---

## 🟨 COLUMN: TODO

### Infrastructure
- Setup Go
- Setup DB
- Setup Docker

### Core
- Transaction engine
- Zone engine
- Pricing engine
- Payment engine

### API
- Auth middleware
- Handlers
- Routes

---

## 🟧 COLUMN: IN PROGRESS

Rules:
- Max 3 cards
- Finish before new task

Example:
- Implement fee calculator
- OCR integration

---

## 🟩 COLUMN: REVIEW

Checklist:
- Code clean
- Test pass
- No panic
- Log ok

---

## 🟪 COLUMN: TESTING

- Unit test
- Flow test
- Load test

---

## ⬜ COLUMN: DONE

Criteria:
- Feature works
- Tested
- Documented

---

## 📋 CARD TEMPLATE

Title:
[CORE] Transaction State Machine

Description:
- Implement state logic
- Add validation
- Add tests

Checklist:
- [ ] Code
- [ ] Test
- [ ] Doc
- [ ] Review

---

## 🏷️ LABEL SYSTEM

CORE      → Red  
API       → Blue  
BUG       → Yellow  
URGENT    → Orange  
TEST      → Green  
DOC       → Purple

---

## 📊 WEEKLY TRACKING

Every week:

- Total cards done: ___
- Bugs found: ___
- Tech debt: ___

