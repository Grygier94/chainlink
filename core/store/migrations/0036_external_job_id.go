package migrations

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"gorm.io/gorm"
)

const (
	up36_1 = `
               ALTER TABLE direct_request_specs DROP COLUMN on_chain_job_spec_id;
               ALTER TABLE webhook_specs DROP COLUMN on_chain_job_spec_id;
               ALTER TABLE vrf_specs ADD CONSTRAINT vrf_specs_public_key_fkey FOREIGN KEY (public_key) REFERENCES encrypted_vrf_keys(public_key) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE;
               ALTER TABLE jobs ADD COLUMN external_job_id uuid; 
	`
	up36_2 = `
               ALTER TABLE jobs 
                    ALTER COLUMN external_job_id SET NOT NULL,
                    ADD CONSTRAINT external_job_id_uniq UNIQUE(external_job_id),
                    ADD CONSTRAINT non_zero_uuid_check CHECK (external_job_id <> '00000000-0000-0000-0000-000000000000');
	`
	down36 = `
               ALTER TABLE direct_request_specs ADD COLUMN on_chain_job_spec_id bytea;
               ALTER TABLE webhook_specs ADD COLUMN on_chain_job_spec_id bytea;
               ALTER TABLE jobs DROP CONSTRAINT external_job_id_uniq;
               ALTER TABLE vrf_specs DROP CONSTRAINT vrf_specs_public_key_fkey;
    `
)

func init() {
	Migrations = append(Migrations, &Migration{
		ID: "0036_external_job_id",
		Migrate: func(db *gorm.DB) error {
			// Add the external ID column and remove type specific ones.
			if err := db.Exec(up36_1).Error; err != nil {
				return err
			}

			// Update all jobs to have an external_job_id.
			// We do this to avoid using the uuid postgres extension.
			var jobs []job.Job
			if err := db.Find(&jobs).Error; err != nil {
				return err
			}
			if len(jobs) != 0 {
				stmt := `UPDATE jobs AS j SET external_job_id = vals.external_job_id FROM (values `
				for i := range jobs {
					if i == len(jobs)-1 {
						stmt += fmt.Sprintf("(uuid('%s'), %d))", uuid.NewV4(), jobs[i].ID)
					} else {
						stmt += fmt.Sprintf("(uuid('%s'), %d),", uuid.NewV4(), jobs[i].ID)
					}
				}
				stmt += ` AS vals(external_job_id, id) WHERE vals.id = j.id`
				if err := db.Exec(stmt).Error; err != nil {
					return err

				}
			}

			// Add constraints on the external_job_id.
			return db.Exec(up36_2).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down36).Error
		},
	})
}
