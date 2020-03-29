CREATE TABLE callws.call_data(
  caller                       text,
  callee                       text NOT NULL,
  start_time                   timestamp NOT NULL,
  end_time                     timestamp,
  inbound                      boolean NOT NULL,
  duration                     integer NOT NULL,
  call_cost                    bigint NOT NULL,
  PRIMARY KEY (caller, callee, start_time)
);

CREATE TABLE callws.call_metadata(
  start_time                   timestamp PRIMARY KEY,
  end_time                     timestamp,
  total_inbound_duration       integer NOT NULL,
  total_outbound_duration      integer NOT NULL,
  total_calls                  integer NOT NULL,
  total_call_costs             integer NOT NULL,
  calls_by_caller              json NOT NULL,
  calls_by_callee              json NOT NULL
);

