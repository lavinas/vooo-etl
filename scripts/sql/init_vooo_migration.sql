-- INSERTING INITIAL DATA
use vooo_migration;

-- _schma
insert into _schma (name, ref_type, init_key) value ('vooo_prod_backend', 'external', -1);
insert into _schma (name, ref_type, init_key) value ('vooo_prod_rawdata', 'internal', 824131765);

-- _expt
insert into _expt (stat) value ("table_name like 'shellbox%'");
insert into _expt (stat) value ("table_name like 'tmp_%'");
insert into _expt (stat) value ("table_name like 'temp_%'");
insert into _expt (stat) value ("table_name like 'analise%'");
insert into _expt (stat) value ("table_name = 'mercadopago_transacao31.03'");
insert into _expt (stat) value ("table_name = 'mercadopago_transacao31.03'");
insert into _expt (stat) value ("table_name = 'connect_item_btg'");
insert into _expt (stat) value ("table_name = 'group_bkp'");
insert into _expt (stat) value ("table_name = 'rel_ant'");
insert into _expt (stat) value ("table_name = 'sales_statement'");
insert into _expt (stat) value ("table_name = 'stix_1405'");
insert into _expt (stat) value ("table_name = 'transacao_btg'");
insert into _expt (stat) value ("table_name = 'transacoes_erradas'");
insert into _expt (stat) value ("table_name = 'transacoes_sap_08_07'");
insert into _expt (stat) value ("table_name = 'brazil_bank_holidays'");
insert into _expt (stat) value ("table_name = 'output_connection_item_back'");
insert into _expt(stat) values ("table_name = 'projections_R_history'");
insert into _expt(stat) value ("table_name = 'pagseguro_control'");
insert into _expt(stat) value ("table_name = 'projections_R'");
insert into _expt(stat) value ("table_name = 'nespresso_converter'");

-- _ref
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client', 'id_aggregator', 'vooo_prod_backend', 'aggregator', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'address', 'id', 'vooo_prod_backend', 'client', 'id_address');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'billing', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'cielo_auth', 'id_client', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'cielo_auth_log', 'id_client', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'cielo_auth_merchants', 'id_client', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_average_sales', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_bank_balance', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_edi_extra_path', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_extra_contract', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_status_detail', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_status_log', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_crawler', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_edi', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_edi_tivit', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_pagador', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_rest', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_sales_system', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_sitef_provider', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_connector_status', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_customization', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_dashboard', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_external_id', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_notprocessing_alert', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_processing_notification', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_provider_composition', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_service', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'client_user', 'id_client', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_crawler', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_edi', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_error', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_item', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_rest', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_sales_lot', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'connection_sales_system', 'connection_id', 'vooo_prod_backend', 'connection', 'connection_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'daily_reports_rates', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'daily_reports_receipts', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'daily_reports_sales', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'document', 'id', 'vooo_prod_backend', 'client', 'id_document');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'files_aws_s3_client', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'files_aws_s3_control', 'client_folder_id', 'vooo_prod_backend', 'files_aws_s3_client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'generic_layout', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'generic_layout_connect_fields', 'layout_id', 'vooo_prod_backend', 'generic_layout', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'generic_layout_map_values', 'field_id', 'vooo_prod_backend', 'generic_layout_connect_fields', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'generic_layout_upload', 'layout_id', 'vooo_prod_backend', 'generic_layout', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'monthly_reports_rates', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'monthly_reports_receipts', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_agenda', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_agenda_reversal', 'output_agenda_id', 'vooo_prod_backend', 'output_agenda', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_agenda_sales', 'output_agenda_id', 'vooo_prod_backend', 'output_agenda', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_agenda_settlement', 'output_agenda_id', 'vooo_prod_backend', 'output_agenda', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_agenda_settlement_reference', 'output_agenda_id', 'vooo_prod_backend', 'output_agenda', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_alert', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_anticipation', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_composition', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_composition_item', 'output_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_connection_item', 'output_item_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_connection_item_back', 'output_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_item', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_item_composition_reference', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_manual_conciliation', 'output_sales_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_outstanding', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_reconciliation', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_reconciliation_alert', 'output_reconciliation_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_reference', 'referrer_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_report_files', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_reversal', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_reversal_adjusts', 'id', 'vooo_prod_backend', 'output_reversal', 'output_reversal_adjusts_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_reversal_bin', 'output_reversal_id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_sales', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_sales_adjusts', 'id', 'vooo_prod_backend', 'output_sales', 'output_sales_adjusts_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_sales_bin', 'output_sales_id', 'vooo_prod_backend', 'output_sales', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_sales_fee', 'output_sales_id', 'vooo_prod_backend', 'output_sales', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_sales_performance', 'id', 'vooo_prod_backend', 'output_sales', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_sales_settlement_item', 'output_sales_id', 'vooo_prod_backend', 'output_sales', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_settlement', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_settlement_bank', 'id', 'vooo_prod_backend', 'output_settlement', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_settlement_files', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_settlement_item', 'output_settlement_id', 'vooo_prod_backend', 'output_settlement', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_statement', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'output_unschedule', 'id', 'vooo_prod_backend', 'output_item', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'pagarme_payable_control', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'pagarme_transaction_control', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'r_admin_fee', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'r_client_exchange_fee', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'r_client_interchange_fee', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'reconciliation_client_provider_actions', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'reconciliation_client_provider_default_account', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'reconciliation_client_rule', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'reconciliation_fee_config', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'reconciliation_transaction_fee', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'rede_request', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'sales', 'item_id', 'vooo_prod_backend', 'connection_item', 'item_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'sales_closed_period', 'item_id', 'vooo_prod_backend', 'connection_item', 'item_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'sales_complement', 'item_id', 'vooo_prod_backend', 'sales', 'item_id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'sitef_control', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user', 'id_client', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user_backoffice', 'id_user', 'vooo_prod_backend', 'user', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user_control', 'user_id', 'vooo_prod_backend', 'user', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user_login_control', 'user_id', 'vooo_prod_backend', 'user', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user_notification_processing', 'user_id', 'vooo_prod_backend', 'user', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user_old_passwords', 'user_id', 'vooo_prod_backend', 'user', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'user_status_log', 'user_id', 'vooo_prod_backend', 'user', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_prod_backend', 'reconciliation_connection_item', 'item_id', 'vooo_prod_backend', 'connection_item', 'item_id');
# cache
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_cache', 'view_agenda', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_cache', 'cash_flow', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_cache', 'view_summary_taxas', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_cache', 'cache_statement', 'client_id', 'vooo_prod_backend', 'client', 'id');
insert into _ref (referrer_schema, referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field) value ('vooo_cache', 'view_summary_agenda', 'client_id', 'vooo_prod_backend', 'client', 'id');

# _key
insert into _key (object, field, init_key) values ('aggregator', 'id', 999999999);
insert into _key (object, field, init_key) values ('billing', 'connection_id', 112706531);
insert into _key (object, field, init_key) values ('connection', 'connection_id', 112705476);
insert into _key (object, field, init_key) values ('connection_crawler', 'connection_id', 112705476);
insert into _key (object, field, init_key) values ('connection_edi', 'connection_id', 112705476);
insert into _key (object, field, init_key) values ('connection_error', 'connection_error_id', 10001155);
insert into _key (object, field, init_key) values ('connection_item', 'item_id', 824131765);
insert into _key (object, field, init_key) values ('connection_rest', 'connection_id', 112705476);
insert into _key (object, field, init_key) values ('connection_sales_lot', 'connection_id', 112705476);
insert into _key (object, field, init_key) values ('connection_sales_system', 'connection_id', 112705476);
insert into _key (object, field, init_key) values ('files_control_action', 'id', 9556265);
insert into _key (object, field, init_key) values ('files_control_vector', 'id', 1370);
insert into _key (object, field, init_key) values ('files_control_vector_ftp', 'id', 1370);
insert into _key (object, field, init_key) values ('output_agenda', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_reversal', 'output_agenda_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_reversal', 'output_reversal_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_sales', 'output_agenda_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_sales', 'output_sales_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_settlement', 'output_agenda_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_settlement', 'output_settlement_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_settlement', 'output_anticipation_id', 510629466);
insert into _key (object, field, init_key) values ('output_agenda_settlement_reference', 'id', 128078527);
insert into _key (object, field, init_key) values ('output_alert', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_anticipation', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_composition', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_connection_item', 'output_item_id', 510629466);
insert into _key (object, field, init_key) values ('output_connection_item', 'connection_item_id', 824131765);
insert into _key (object, field, init_key) values ('output_connection_item_back', 'output_item_id', 510629466);
insert into _key (object, field, init_key) values ('output_connection_item_back', 'connection_item_id', 824131765);
insert into _key (object, field, init_key) values ('output_item', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_item_composition_reference', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_manual_conciliation', 'output_sales_id', 510629466);
insert into _key (object, field, init_key) values ('output_outstanding', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_reconciliation', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_reconciliation_alert', 'output_reconciliation_id', 510629466);
insert into _key (object, field, init_key) values ('output_reconciliation_alert', 'output_alert_id', -1);
insert into _key (object, field, init_key) values ('output_reference', 'referenced_id', 510629466);
insert into _key (object, field, init_key) values ('output_reference', 'referrer_id', 510629466);
insert into _key (object, field, init_key) values ('output_reversal', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_reversal_adjusts', 'id', 576137);
insert into _key (object, field, init_key) values ('output_reversal_bin', 'output_reversal_id', 510629466);
insert into _key (object, field, init_key) values ('output_sales', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_sales_adjusts', 'id', 7992112);
insert into _key (object, field, init_key) values ('output_sales_bin', 'output_sales_id', 510629466);
insert into _key (object, field, init_key) values ('output_sales_fee', 'id', 232056358);
insert into _key (object, field, init_key) values ('output_sales_performance', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_sales_settlement_item', 'output_sales_id', 510629466);
insert into _key (object, field, init_key) values ('output_sales_settlement_item', 'output_settlement_item_id', 319343061);
insert into _key (object, field, init_key) values ('output_settlement', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_settlement_bank', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_settlement_item', 'id', 319343061);
insert into _key (object, field, init_key) values ('output_statement', 'id', 510629466);
insert into _key (object, field, init_key) values ('output_unschedule', 'id', 510629466);
insert into _key (object, field, init_key) values ('rede_files_action', 'id', 7028);
insert into _key (object, field, init_key) values ('rede_files_control', 'id', 2774);
insert into _key (object, field, init_key) values ('rede_request', 'id', 86);
insert into _key (object, field, init_key) values ('sales', 'item_id', 824131765);
insert into _key (object, field, init_key) values ('sales_cancel_log', 'id', 529);
insert into _key (object, field, init_key) values ('sales_closed_period', 'item_id', 824131765);
insert into _key (object, field, init_key) values ('sales_complement', 'item_id', 824131765);
insert into _key (object, field, init_key) values ('sitef_control', 'id', 1012);
insert into _key (object, field, init_key) values ('ticket_files_action', 'id', 4729);
insert into _key (object, field, init_key) values ('ticket_files_control', 'id', 27792);
insert into _key (object, field, init_key) values ('tivit_control', 'id', 58);
insert into _key (object, field, init_key) values ('user_login_control', 'id', 146647);
insert into _key (object, field, init_key) values ('user_old_passwords', 'id', 275);
insert into _key (object, field, init_key) values ('user_status_log', 'id', 57);

commit;

### delete keys + update
# output_reconciliation_alert.output_alert.id
# output_connection_item.connection_item_id
# output_reference.referrer_id
# daily_reports_rates.day
# daily_reports_rates.month
# daily_reports_rates.payment_method_id
# daily_reports_rates.provider_id
# daily_reports_rates.year
# billing.aggregator
# billing.client_id
# billing.processing_date
# billing.provider_id
# billing.type_transactions

# all
# vooo_prod_backend.purge_backend_table_list
# vooo_prod_backend.purge_rawdata_table_list
# vooo_prod_backend.generic_layout_map_values
# billing.billing_control

# none
# vooo_prod_backend.aggregator


## billing
# history
INSERT INTO job values (2002,'vooo_billing.billing_history','table','copy','vooo_billing','billing_history',0);
INSERT INTO job_key VALUES (200200,2002,'id',154,5000);
insert into ref values (200200, 2002, 1);
insert into ref_key values (200200, 2002, 'id_aggregator', 'id');
# amount
INSERT INTO job VALUES (2003,'vooo_billing.billing_amount','table','copy','vooo_billing','billing_amount',0);
INSERT INTO job_key VALUES (200300,2003,'id',35326,5000);
insert into ref values (200300, 2003, 2);
insert into ref_key values (200300, 2003, 'client_id', 'id');


# others
insert into job values (837, 'vooo_prod_backend.output_settlement_item-2', 'table', 'copy', 'vooo_prod_backend', 'output_settlement_item', 0);
insert into job_key values (83700, 837, 'id', 319343061, 10000);
insert into ref values (83700, 837, 49);
insert into ref_key values (83700, 83700, 'output_anticipation_id', 'id');