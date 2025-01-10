// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a pgEdge database.
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'pgedge:index/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * Backup configuration for the database.
     */
    public readonly backups!: pulumi.Output<outputs.DatabaseBackups>;
    /**
     * The ID of the cluster this database belongs to.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * List of components in the database.
     */
    public /*out*/ readonly components!: pulumi.Output<outputs.DatabaseComponent[]>;
    /**
     * The configuration version of the database.
     */
    public readonly configVersion!: pulumi.Output<string>;
    /**
     * The timestamp when the database was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Display name for the database. Maximum length is 25 characters.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The domain associated with the database.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * Extensions configuration for the database.
     */
    public readonly extensions!: pulumi.Output<outputs.DatabaseExtensions>;
    /**
     * The name of the database.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Map of nodes in the database.
     */
    public readonly nodes!: pulumi.Output<{[key: string]: outputs.DatabaseNodes}>;
    /**
     * A list of options for the database.
     */
    public readonly options!: pulumi.Output<string[] | undefined>;
    /**
     * The PostgreSQL version of the database.
     */
    public /*out*/ readonly pgVersion!: pulumi.Output<string>;
    /**
     * List of roles in the database.
     */
    public readonly roles!: pulumi.Output<outputs.DatabaseRole[]>;
    /**
     * The current status of the database.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            resourceInputs["backups"] = state ? state.backups : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["components"] = state ? state.components : undefined;
            resourceInputs["configVersion"] = state ? state.configVersion : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["extensions"] = state ? state.extensions : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["pgVersion"] = state ? state.pgVersion : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.nodes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodes'");
            }
            resourceInputs["backups"] = args ? args.backups : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["configVersion"] = args ? args.configVersion : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["extensions"] = args ? args.extensions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodes"] = args ? args.nodes : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["components"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["pgVersion"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * Backup configuration for the database.
     */
    backups?: pulumi.Input<inputs.DatabaseBackups>;
    /**
     * The ID of the cluster this database belongs to.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * List of components in the database.
     */
    components?: pulumi.Input<pulumi.Input<inputs.DatabaseComponent>[]>;
    /**
     * The configuration version of the database.
     */
    configVersion?: pulumi.Input<string>;
    /**
     * The timestamp when the database was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Display name for the database. Maximum length is 25 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The domain associated with the database.
     */
    domain?: pulumi.Input<string>;
    /**
     * Extensions configuration for the database.
     */
    extensions?: pulumi.Input<inputs.DatabaseExtensions>;
    /**
     * The name of the database.
     */
    name?: pulumi.Input<string>;
    /**
     * Map of nodes in the database.
     */
    nodes?: pulumi.Input<{[key: string]: pulumi.Input<inputs.DatabaseNodes>}>;
    /**
     * A list of options for the database.
     */
    options?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PostgreSQL version of the database.
     */
    pgVersion?: pulumi.Input<string>;
    /**
     * List of roles in the database.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.DatabaseRole>[]>;
    /**
     * The current status of the database.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Backup configuration for the database.
     */
    backups?: pulumi.Input<inputs.DatabaseBackups>;
    /**
     * The ID of the cluster this database belongs to.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The configuration version of the database.
     */
    configVersion?: pulumi.Input<string>;
    /**
     * Display name for the database. Maximum length is 25 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Extensions configuration for the database.
     */
    extensions?: pulumi.Input<inputs.DatabaseExtensions>;
    /**
     * The name of the database.
     */
    name?: pulumi.Input<string>;
    /**
     * Map of nodes in the database.
     */
    nodes: pulumi.Input<{[key: string]: pulumi.Input<inputs.DatabaseNodes>}>;
    /**
     * A list of options for the database.
     */
    options?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of roles in the database.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.DatabaseRole>[]>;
}
