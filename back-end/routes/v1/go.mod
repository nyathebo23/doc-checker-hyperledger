module feuli/backend/routes

go 1.18

replace feuli/backend/routes/v1/document_saves => ./document_saves

replace feuli/backend/routes/v1/organizations => ./organizations

replace feuli/backend/routes/v1/beneficiaries => ./beneficiaries

replace feuli/backend/models => ../../models

replace feuli/backend/hyperledger => ../../hyperledger
